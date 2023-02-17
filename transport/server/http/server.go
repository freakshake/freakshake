package httpserver

import (
	"fmt"
	"net/http"
)

func Listen(h http.Handler, errCh chan error, ip, port string) *http.Server {
	addr := ip + ":" + port

	srv := &http.Server{
		Addr:    addr,
		Handler: h,
	}

	go func() {
		fmt.Println("[+] HTTP server listening on", addr)

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errCh <- err
		}
	}()

	return srv
}
