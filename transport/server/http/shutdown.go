package httpserver

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func SpawnShutdownListener(errCh chan error, httpServer *http.Server) {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	go func() {
		<-ctx.Done()

		fmt.Println("\n[-] shutdown signal received")

		httpServer.SetKeepAlivesEnabled(false)

		ctxTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		if err := httpServer.Shutdown(ctxTimeout); err != nil {
			errCh <- err
		}

		defer func() {
			stop()
			cancel()
			close(errCh)
		}()

		fmt.Println("[-] shutdown completed")
	}()
}
