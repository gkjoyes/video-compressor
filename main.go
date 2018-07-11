package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/g-kutty/v-comp/compressor"

	"github.com/g-kutty/go-streamer/graceful"
	"github.com/g-kutty/v-comp/logger"
	"github.com/g-kutty/v-comp/watcher"
)

const appVersion = "1.0"

var (
	path    string
	help    bool
	version bool
)

// Read command line arguments before start.
func init() {
	flag.StringVar(&path, "p", "", "")
	flag.StringVar(&path, "path", "", "")
	flag.BoolVar(&version, "v", false, "")
	flag.BoolVar(&version, "version", false, "")
	flag.BoolVar(&help, "h", false, "")
	flag.BoolVar(&help, "help", false, "")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: v-comp [options]\n")
		fmt.Fprintf(os.Stderr, "options:\n")
		fmt.Fprintf(os.Stderr, "\t-p, -path      Directory                The directory to watch.\n")
		fmt.Fprintf(os.Stderr, "\t-v, -version   Version                  Prints the version.\n")
		fmt.Fprintf(os.Stderr, "\t-h, -help      Help                     Show this help.\n")
	}
}

func main() {
	go graceful.ActivateGracefulShutdown()
	parseFlags()

	var files []string

	// path
	root := "/home/kutty/Videos"

	// walk through files.
	err := filepath.Walk(root, watcher.Visit(&files))
	if err != nil {
		log.Fatal(err)
	}

	// compress all video files.
	res := compressor.Compress(files, 4)
	fmt.Println("---------got response-----------", res)
}

// parseFlags read command line arguments.
func parseFlags() {
	flag.Parse()

	// display version.
	if version {
		fmt.Printf("v-comp v%s\n", appVersion)
		os.Exit(0)
	}

	// display help guides for command line arguments.
	if help {
		flag.Usage()
		os.Exit(0)
	}
	validateFlags()
}

// validateFlags validate the flag values.
func validateFlags() {
	if path == "" {
		// Get the current working directory if no path is specified.
		dir, _ := os.Getwd()
		path, _ = filepath.Abs(dir)
	} else {
		dir, err := os.Stat(path)
		if err != nil {
			logger.Error().Message("Cannot find path,", logger.FormattedMessage(path)).Log()
			os.Exit(1)
		}
		if !dir.IsDir() {
			logger.Error().Message(fmt.Sprintf("Invalid path, %s. Path must be directory.", logger.FormattedMessage(path))).Log()
			os.Exit(1)
		}
		path, _ = filepath.Abs(path)
	}
}
