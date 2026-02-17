package main

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/DumbNoxx/goxe/pkg/pipelines"
)

// executeHandoff coordinates the graceful shutdown of log producers and processors.
//
// Parameters:
//
//   - once: sync.Once to ensure the handoff logic is executed only once.
//   - cancel: cancellation function for the main context, used to stop producers.
//   - pipe: log input channel (will be closed at the end).
//   - wgProc: WaitGroup for processors (waits for them to finish).
//   - wgProd: WaitGroup for producers (waits with a defined timeout).
//
// Returns:
//
//   - void: the function blocks until all processors have finished execution.
//
// The function performs:
//
//   - Executes the logic exactly once using once.Do.
//   - Calls cancel() to signal producers that they must stop.
//   - Launches a goroutine that waits for wgProd to reach zero (producers finishing).
//   - Waits up to 2 seconds for that goroutine to complete; if the timeout expires,
//     it prints "[System] Handoff: Force closing producers..." and proceeds.
//   - Closes the pipe channel so processors know no more data is coming.
//   - Waits for wgProc (processors) to finish using wgProc.Wait().
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

// handleUpdate manages system signals for updates or interruptions.
//
// Parameters:
//
//   - sigChan: channel for receiving operating system signals.
//   - ctx: main context for cancellation.
//   - cancel: function to cancel the main context.
//   - pipe: log input channel.
//   - wgProcessor: WaitGroup for processors.
//   - wgProducer: WaitGroup for producers.
//   - once: sync.Once to ensure the handoff logic is executed only once.
//
// Returns:
//
//   - void: the function runs in a loop until an update signal is received
//     or the context is cancelled.
//
// The function performs:
//
//   - Continuously listens for signals on the sigChan channel.
//
//   - For each received signal:
//
//     -If it is an update signal (determined by isUpdateSignal):
//
//     1. Prints an update initiation message.
//
//     2. Creates a ticker that ticks every second.
//
//     3. Counts up to 5, printing the countdown.
//
//     4. Upon reaching 5, it calls autoUpdate with the parameters and terminates.
//
//     -If the signal is os.Interrupt (Ctrl+C), it calls cancel() to stop the context.
//
//   - If the context is cancelled during the countdown, the function returns immediately.
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
