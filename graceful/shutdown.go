package graceful

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/g-kutty/v-comp/logger"
)

// ActivateGracefulShutdown handle graceful shutdown.
func ActivateGracefulShutdown() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	for sig := range signalChan {
		if sig == syscall.SIGINT {
			logger.Warn().Command("interrupt", "i").Message("Exiting.").Log()
			os.Exit(0)
		}
	}
}
