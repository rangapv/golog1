package main

import (
	"log"

	"github.com/rangapv/golog1/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
