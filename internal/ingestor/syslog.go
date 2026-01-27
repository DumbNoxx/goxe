package ingestor

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
	"unsafe"

	"github.com/DumbNoxx/Goxe/internal/options"
	"github.com/DumbNoxx/Goxe/pkg/pipelines"
)

var (
	PORT      string = ":" + strconv.Itoa(options.Config.Port)
	lastIp    string
	lastRawIp net.IP
)

func Udp(pipe chan<- *pipelines.LogEntry, wg *sync.WaitGroup) {

	addr, err := net.ResolveUDPAddr("udp", PORT)
	if err != nil {
		fmt.Println("Error resolving address:", err)
		return
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Listening error:", err)
		return
	}
	defer conn.Close()

	fmt.Printf("Server listening on port %s\n", PORT)

	for {
		buffer := pipelines.BufferPool.Get().([]byte)
		n, clientAddr, err := conn.ReadFromUDP(buffer)

		if err != nil {
			fmt.Println("Error reading", err)
			pipelines.BufferPool.Put(buffer)
			continue
		}

		if !clientAddr.IP.Equal(lastRawIp) {
			lastRawIp = clientAddr.IP
			lastIp = clientAddr.IP.String()
		}

		dates := pipelines.EntryPool.Get().(*pipelines.LogEntry)
		dates.Content = unsafe.String(&buffer[0], n)
		dates.Source = lastIp
		dates.Timestamp = time.Now()
		dates.IdLog = options.Config.IdLog
		dates.RawEntry = buffer

		pipe <- dates
	}

}
