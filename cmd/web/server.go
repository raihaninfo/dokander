package main

import (
	"fmt"
	"net/http"
	"time"
)

func (a *application) ListenAndServe() error {
	url := fmt.Sprintf("%s:%s", a.Server.Host, a.Server.Port)
	srv := http.Server{
		Handler:     a.router(),
		Addr:        url,
		ReadTimeout: 300 * time.Second,
	}
	return srv.ListenAndServe()
}
