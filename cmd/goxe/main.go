package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"

	"github.com/DumbNoxx/Goxe/internal/ingestor"
	"github.com/DumbNoxx/Goxe/internal/processor"
	"github.com/DumbNoxx/Goxe/pkg/pipelines"
)

func main() {
	var wg sync.WaitGroup
	pipe := make(chan pipelines.LogEntry)
	var idLog string
	var mu sync.Mutex

	fmt.Println("Antes de continuar, por favor ¿Puedes proporcionar el IdLog de tu servidor?")
	fmt.Println("Si tu servidor no tiene id log escribe 9")
	for {
		fmt.Scan(&idLog)
		n, _ := strconv.ParseInt(idLog, 10, 64)
		if n == 9 {
			idLog = ""
			break
		}
		if n != 9 && n != 0 {
			fmt.Println("Debes de introducir el 9")
			continue
		}

		if len(idLog) < 4 {
			fmt.Println("Tu idLog debe de ser mayor a 4 caracteres, si tu servidor no tiene id log escribe 9")
			continue
		}
		break
	}

	wg.Add(1)
	go processor.Clean(pipe, &wg, &mu)
	go ingestor.Udp(pipe, &wg, idLog)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	close(pipe)

	wg.Wait()

}
