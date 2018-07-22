package compressor

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/g-kutty/v-comp/logger"
)

// file define each file
type file struct {
	name string
	idx  int
}

// Compress video files concurrently.
func Compress(files []string, threads int) error {
	var wg sync.WaitGroup

	// make sure that all called resources should release even if no errors.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// errors channel.
	errs := make(chan error, threads)

	// videos channel.
	fl := len(files)
	var ch = make(chan file, fl)

	// this starts threads number of goroutines that wait for something to do.
	wg.Add(threads)
	for i := 0; i < threads; i++ {
		go func(n int) {

			// check if any error occurred in any other goroutines.
			select {
			// error somewhere, terminate.
			case <-ctx.Done():
				return
			// default is must to avoid blocking.
			default:
			}

			for {
				f, ok := <-ch
				if !ok {
					wg.Done()
					return
				}

				// compress each videos.
				{
					filePath := strings.Split(f.name, "/")
					l := len(filePath) - 1

					fileName := filePath[l]
					newFilePath := strings.Join(filePath[:l], "/") + "/compress/"

					// create compress folder if not exists.
					if _, err := os.Stat(newFilePath); os.IsNotExist(err) {
						err = os.MkdirAll(newFilePath, 0755)
						if err != nil {
							errs <- err
							cancel()
							return
						}
					}

					// log
					logger.Info().Command("compress", "c").Message(fmt.Sprintf("[%d/%d]", f.idx, fl) + logger.FormattedMessage(newFilePath+fileName)).Log()

					// compress video using ffmpeg.
					if _, err := os.Stat(newFilePath + fileName); os.IsNotExist(err) {
						_, err := exec.Command("ffmpeg", "-i", f.name, "-acodec", "mp3", newFilePath+fileName, "-y").Output()
						if err != nil {
							errs <- err
							cancel()
							return
						}
					}
				}
			}
		}(i)
	}

	// now the jobs can be added to the channel, which is used as a queue.
	for i, v := range files {
		ch <- file{name: v, idx: i}
	}

	close(ch)
	wg.Wait()

	// return error, if any.
	if ctx.Err() != nil {
		return <-errs
	}
	return nil
}
