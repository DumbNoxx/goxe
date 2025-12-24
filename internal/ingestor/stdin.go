package ingestor

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/DumbNoxx/Goxe/internal/pipelines"
)

// Function to read the received information
func IngestorData(pipe chan<- pipelines.LogEntry, wg *sync.WaitGroup) {
	defer wg.Done()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("enter text (ctrl + d or ctrl + c to exit)")

	for scanner.Scan() {
		line := scanner.Text()
		log := pipelines.LogEntry{
			Content:   line,
			Source:    "STDIN",
			Timestamp: time.Now(),
		}
		pipe <- log
	}

	if err := scanner.Err(); err != nil {
		if err != os.ErrClosed {
			fmt.Fprintf(os.Stderr, "error reading stdin: %v\n", err)
		}
	}
}
