package ingestor

import (
	"bufio"
	"fmt"
	"github.com/DumbNoxx/Goxe/internal/processor"
	"os"
	"sync"
)

func IngestorData() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("enter text (ctrl + d or ctrl + c to exit)")
	pipe := make(chan string)
	var wg sync.WaitGroup

	wg.Add(1)
	go processor.Clean(pipe, &wg)

	for scanner.Scan() {
		line := scanner.Text()
		pipe <- line
	}

	if err := scanner.Err(); err != nil {
		if err != os.ErrClosed {
			fmt.Fprintf(os.Stderr, "error reading stdin: %v\n", err)
		}
	}
	close(pipe)
	wg.Wait()
}
