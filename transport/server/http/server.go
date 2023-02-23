package httpserver

import (
	"fmt"
	"net/http"
)

func SpawnListener(srv *http.Server, errCh chan error) {
	go func() {
		fmt.Println("[+] HTTP server listening on", srv.Addr)

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errCh <- err
		}
	}()
}
