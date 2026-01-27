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
	entryPool        = sync.Pool{
		New: func() any { return new(pipelines.LogEntry) },
	}
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

	buf := sync.Pool{
		New: func() any {
			return make([]byte, 1024)

		},
	}

	for {
		buffer := buf.Get().([]byte)
		n, clientAddr, err := conn.ReadFromUDP(buffer)

		if err != nil {
			fmt.Println("Error reading", err)
			buf.Put(buffer)
			continue
		}

		if !clientAddr.IP.Equal(lastRawIp) {
			lastRawIp = clientAddr.IP
			lastIp = clientAddr.IP.String()
		}

		dates := entryPool.Get().(*pipelines.LogEntry)
		dates.Content = unsafe.String(&buffer[0], n)
		dates.Source = lastIp
		dates.Timestamp = time.Now()
		dates.IdLog = options.Config.IdLog

		pipe <- dates
	}

}
