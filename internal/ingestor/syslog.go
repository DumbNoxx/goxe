package ingestor

import (
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/DumbNoxx/Goxe/internal/pipelines"
)

var (
	PORT = ":5642"
)

func Udp(pipe chan<- pipelines.LogEntry, wg *sync.WaitGroup) {
	defer wg.Done()

	addr, err := net.ResolveUDPAddr("udp", PORT)
	if err != nil {
		fmt.Println("Error al resolver la dirección", err)
		return
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error al escuchar: ", err)
		return
	}
	defer conn.Close()

	fmt.Printf("Servidor escuchando en http://localhost%s\n", PORT)

	buffer := make([]byte, 1024)

	for {
		n, clientAddr, err := conn.ReadFromUDP(buffer)

		if err != nil {
			fmt.Println("error al leer", err)
			continue
		}

		message := string(buffer[:n])

		dates := pipelines.LogEntry{
			Source:    clientAddr.String(),
			Content:   message,
			Timestamp: time.Now(),
		}

		pipe <- dates
	}

}
