package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"slices"
	"sync"

	"github.com/DumbNoxx/goxe/internal/ingestor"
	"github.com/DumbNoxx/goxe/internal/options"
	"github.com/DumbNoxx/goxe/internal/processor"
	"github.com/DumbNoxx/goxe/pkg/pipelines"
)

var (
	versionFlag *bool
	isUpgrade   *bool
	version     string
)

func init() {
	versionFlag = flag.Bool("v", false, "")
	isUpgrade = flag.Bool("is-upgrade", false, "Internal use for hot-swap")
}

func main() {
	flag.Parse()
	arg := os.Args

	if slices.Contains(arg, "update") {
		fmt.Println("Sending update signal to the active instance...")
		cmd := exec.Command("pkill", "-SIGUSR1", "goxe")
		cmd.Run()
		return
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	stopChan := make(chan os.Signal, 1)
	pipe := make(chan *pipelines.LogEntry, 100)
	var (
		wgProcessor sync.WaitGroup
		wgProducer  sync.WaitGroup
		mu          sync.Mutex
		once        sync.Once
	)

	if *versionFlag {
		fmt.Println(getVersion())
		os.Exit(0)
	}
	if *isUpgrade {
		fmt.Println("[System] System updated")
	}

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
