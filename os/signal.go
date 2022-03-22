package os

import (
	"os"
	"os/signal"
)

func WaitInterruptSignal() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
