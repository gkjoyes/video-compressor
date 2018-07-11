package compressor

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Compress video files concurrently.
func Compress(files []string, threads int) int {
	var ch = make(chan string, 50)
	var wg sync.WaitGroup

	// make sure it's called to release resources even if no errors.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// errors channel
	errs := make(chan int, 4)

	// this starts threads number of goroutines that wait for something to do.
	wg.Add(threads)
	for i := 0; i < threads; i++ {
		go func() {
			defer wg.Done()
			// Check if any error occurred in any other gorouties:
			select {
			case <-ctx.Done():
				return // Error somewhere, terminate
			default: // Default is must to avoid blocking
			}
			for {
				file, ok := <-ch
				if !ok {
					wg.Done()
					return
				}
				if file == "a" {
					errs <- 1
					cancel()
					return
				}
				// compress.
				do(file)
			}
		}()
	}

	// now the jobs can be added to the channel, which is used as a queue.
	for _, v := range files {
		ch <- v
	}

	for i := 0; i < 50; i++ {
		ch <- "a"
	}

	// return error, if any.
	if ctx.Err() != nil {
		fmt.Println("----------error------", <-errs)
		return <-errs
	}

	close(ch)
	wg.Wait()
	return 0
}

// compress
func do(file string) {
	fmt.Println("-------file----", file)
	time.Sleep(1000000)
}
