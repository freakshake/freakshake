package httpserver

import (
	"fmt"
	"net/http"
	"time"
)

func Listen(h http.Handler, errCh chan error, ip, port string) *http.Server {
	addr := ip + ":" + port

	srv := &http.Server{
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 3 * time.Second,
		Addr:              addr,
		Handler:           h,
	}

	go func() {
		fmt.Println("[+] HTTP server listening on", addr)

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errCh <- err
		}
	}()

	return srv
}
