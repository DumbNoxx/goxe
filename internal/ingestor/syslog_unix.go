//go:build linux || darwin

package ingestor

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"sync"
	"syscall"
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

func Udp(ctx context.Context, pipe chan<- *pipelines.LogEntry, wg *sync.WaitGroup) {
	defer wg.Done()
	Server.Control = func(networ, addr string, c syscall.RawConn) error {
		err := c.Control(func(fd uintptr) {
			syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1)
		})
		if err != nil {
			return err
		}
		return nil
	}
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
