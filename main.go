package main

import (
	"log"
	"net/http"
	"time"

	"whapi/auth"
	"whapi/router"
)

func main() {
	rootHandler := http.HandlerFunc(router.RootHdlr)

	mux := http.NewServeMux()
	mux.Handle("/", auth.Auth(rootHandler)) //we use general route

	//server
	srv := &http.Server{
		Addr:           ":7000",
		Handler:        mux,
		ReadTimeout:    300 * time.Second,
		WriteTimeout:   300 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
