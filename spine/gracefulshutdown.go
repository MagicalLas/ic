package spine

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var C, Cancel = context.WithCancelCause(context.Background())

var SystemGroup sync.WaitGroup

func init() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-stop
		Cancel(nil)
	}()
}
