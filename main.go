package main

import (
	"fmt"
	"net/http"

	"github.com/chins/go-docker-visualizing/pkg"
	"github.com/chins/go-docker-visualizing/routers"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", pkg.HttpPort),
		Handler:        router,
		ReadTimeout:    pkg.ReadTimeout,
		WriteTimeout:   pkg.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
