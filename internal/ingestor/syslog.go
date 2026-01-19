package ingestor

import (
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/DumbNoxx/Goxe/pkg/pipelines"
)

const (
	PORT string = ":5642"
)

func Udp(pipe chan<- pipelines.LogEntry, wg *sync.WaitGroup, idlog string) {

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

	buffer := make([]byte, 1024)

	for {
		n, clientAddr, err := conn.ReadFromUDP(buffer)

		if err != nil {
			fmt.Println("Error reading", err)
			continue
		}

		message := string(buffer[:n])
		host, _, _ := net.SplitHostPort(clientAddr.String())

		dates := pipelines.LogEntry{
			Source:    host,
			Content:   message,
			Timestamp: time.Now(),
			IdLog:     idlog,
		}

		pipe <- dates
	}

}
