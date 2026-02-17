package main

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/DumbNoxx/goxe/pkg/pipelines"
)

func executeHandoff(once *sync.Once, cancel context.CancelFunc, pipe chan<- *pipelines.LogEntry, wgProc, wgProd *sync.WaitGroup) {
	once.Do(func() {
		cancel()

		done := make(chan struct{})
		go func() {
			wgProd.Wait()
			close(done)
		}()

		select {
		case <-done:
		case <-time.After(2 * time.Second):
			fmt.Println("[System] Handoff: Force closing producers...")
		}

		close(pipe)
		wgProc.Wait()
	})
}

func handleUpdate(sigChan chan os.Signal, ctx context.Context, cancel context.CancelFunc, pipe chan<- *pipelines.LogEntry, wgProcessor, wgProducer *sync.WaitGroup, once *sync.Once) {
	for sig := range sigChan {
		if isUpdateSignal(sig) {

			fmt.Println("\n[System] Update signal received! Starting auto-update...")
			ticker := time.NewTicker(1 * time.Second)
			count := 1
			updateDone := false
			defer ticker.Stop()
			for !updateDone {
				select {
				case <-ticker.C:
					if count != 5 {
						fmt.Printf("%d..", count)
					}
					if count == 5 {
						fmt.Printf("%d\n", count)
						fmt.Println("Updating...")
						autoUpdate(ctx, cancel, pipe, wgProcessor, wgProducer, once)
						updateDone = true
					}
					count++
				case <-ctx.Done():
					return
				}

			}
			return
		}
		if sig == os.Interrupt {
			cancel()
		}
	}
}
