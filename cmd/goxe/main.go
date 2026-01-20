package main

import (
	"os"
	"os/signal"
	"sync"

	"github.com/DumbNoxx/Goxe/internal/ingestor"
	"github.com/DumbNoxx/Goxe/internal/options"
	"github.com/DumbNoxx/Goxe/internal/processor"
	"github.com/DumbNoxx/Goxe/pkg/pipelines"
)

func main() {
	var wg sync.WaitGroup
	pipe := make(chan pipelines.LogEntry)
	var mu sync.Mutex

	options.CacheDirGenerate()

	wg.Add(1)
	go processor.Clean(pipe, &wg, &mu)
	go ingestor.Udp(pipe, &wg)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	close(pipe)

	wg.Wait()

}
