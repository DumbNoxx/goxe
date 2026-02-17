//go:build windows

package ingestor

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/DumbNoxx/goxe/internal/options"
	"github.com/DumbNoxx/goxe/pkg/pipelines"
)

var (
	PORT      string = ":" + strconv.Itoa(options.Config.Port)
	lastIp    string
	lastRawIp net.IP
	Server    net.ListenConfig
)

// Udp listens on a UDP port, receives datagrams, converts them into LogEntry, and sends them to a channel.
// (Windows-specific implementation).
//
// Parameters:
//
//   - ctx: context for cancellation; when cancelled, the function closes the connection and terminates.
//   - pipe: write-only channel (chan<- *pipelines.LogEntry) where log entries will be sent.
//   - wg: pointer to sync.WaitGroup to signal the caller that the goroutine has finished (calls wg.Done() at the end).
//
// Returns:
//
//   - void: the function runs until ctx is cancelled or a critical error occurs.
//
// The function performs:
//
//   - Listens on the port defined by the PORT global constant using Server.ListenPacket.
//   - Adjusts the receive buffer size based on options.Config.BufferUdpSize (in MB).
//   - In a separate goroutine, closes the connection when ctx.Done() is triggered.
//   - In a loop, retrieves a buffer from the pool (pipelines.BufferPool), reads a datagram, and extracts the client's IP address.
//   - Constructs a pipelines.LogEntry with the content, source (IP), timestamp, IdLog (from options.Config), and the original buffer.
//   - Sends the entry to the pipe channel, respecting context cancellation.
//   - If a read error occurs that is not related to cancellation, it logs the error and continues.
//   - Upon termination (via cancellation), the function returns.
func Udp(ctx context.Context, pipe chan<- *pipelines.LogEntry, wg *sync.WaitGroup) {
	defer wg.Done()
	conn, err := Server.ListenPacket(ctx, "udp", PORT)

	if udpConn, ok := conn.(*net.UDPConn); ok {
		udpConn.SetReadBuffer(options.Config.BufferUdpSize * 1024 * 1024)
	}

	if err != nil {
		fmt.Println("Listening error:", err)
		return
	}
	go func() {
		<-ctx.Done()
		conn.Close()
	}()

	fmt.Printf("Goxe listening on port %s\n", PORT)

	for {
		buffer := pipelines.BufferPool.Get().([]byte)
		n, clientAddr, err := conn.ReadFrom(buffer)

		if err != nil {
			if ctx.Err() != nil {
				return
			}
			fmt.Println("Error reading", err)
			pipelines.BufferPool.Put(buffer)
			continue
		}

		if udpAddr, ok := clientAddr.(*net.UDPAddr); ok {
			if !udpAddr.IP.Equal(lastRawIp) {
				lastRawIp = udpAddr.IP
				lastIp = udpAddr.IP.String()
			}
		}

		dates := pipelines.EntryPool.Get().(*pipelines.LogEntry)
		dates.Content = buffer[:n]
		dates.Source = lastIp
		dates.Timestamp = time.Now()
		dates.IdLog = options.Config.IdLog
		dates.RawEntry = buffer

		select {
		case pipe <- dates:
		case <-ctx.Done():
			return
		}

	}

}
