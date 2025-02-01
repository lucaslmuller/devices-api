package shutdown

import (
	"os"
	"os/signal"

	"github.com/lucaslmuller/technical-test/internal/infrastructure"
)

func Gracefully(res *infrastructure.Resources) {
	quitSignal := Subscribe()

	receivedSignal := <-quitSignal

	res.Logger.Info("received a quit signal \"%v\", starting shutdown", receivedSignal)

	closeHTTPServer(res)
	closeConnections(res)

	res.Logger.Info("application shutdown!")
}

func Subscribe() chan os.Signal {
	quitSignal := make(chan os.Signal, 1)
	signal.Notify(quitSignal, os.Interrupt)
	return quitSignal
}
