package main

import (
	"sync"

	"github.com/DumbNoxx/Goxe/internal/ingestor"
	"github.com/DumbNoxx/Goxe/internal/pipelines"
	"github.com/DumbNoxx/Goxe/internal/processor"
)

func main() {
	var wg sync.WaitGroup
	pipe := make(chan pipelines.LogEntry)

	wg.Add(3)
	go processor.Clean(pipe, &wg)
	go ingestor.IngestorData(pipe, &wg)
	go ingestor.Udp(pipe, &wg)

	wg.Wait()
	close(pipe)

}
