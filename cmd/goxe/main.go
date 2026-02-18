package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"slices"
	"sync"

	"github.com/DumbNoxx/goxe/internal/ingestor"
	"github.com/DumbNoxx/goxe/internal/options"
	"github.com/DumbNoxx/goxe/internal/processor"
	"github.com/DumbNoxx/goxe/pkg/pipelines"
)

func main() {
	flag.Parse()
	arg := os.Args
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	switch {
	case slices.Contains(arg, "update"):
		var (
			req http.Request
			res http.Response
		)
		if getVersion() == getVersionLatest(&req, &res, ctx).Tag_name {
			fmt.Println("[System] System is already up to date")
			os.Exit(0)
		}
		updateArg()
		os.Exit(0)
	case versionFlag:
		fmt.Println(getVersion())
		os.Exit(0)
	case *isUpgrade:
		fmt.Println("[System] Goxe updated")
	}

	var (
		stopChan    = make(chan os.Signal, 1)
		pipe        = make(chan *pipelines.LogEntry, 100)
		wgProcessor sync.WaitGroup
		wgProducer  sync.WaitGroup
		mu          sync.Mutex
		once        sync.Once
	)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, watchSignals...)

	go handleUpdate(sigChan, ctx, cancel, pipe, &wgProcessor, &wgProducer, &once)
	options.CacheDirGenerate()

	wgProcessor.Add(1)
	go processor.Clean(ctx, pipe, &wgProcessor, &mu)
	wgProducer.Add(1)
	go ingestor.Udp(ctx, pipe, &wgProducer)
	wgProducer.Add(1)
	go viewConfig(ctx, &wgProducer)
	wgProducer.Add(1)
	go viewNewVersion(ctx, &wgProducer)

	signal.Notify(stopChan, os.Interrupt)
	<-stopChan

	fmt.Println("[System] Shutdown app, flushing buffers...")
	executeHandoff(&once, cancel, pipe, &wgProcessor, &wgProducer)

}
