package main

import (
	"fmt"
	"greact/server"
	"log"
	"net/http"
	"path"
)

var (
	addr      = "localhost:8080"
	buildPath = path.Clean("ui/build")
	buildURL  = fmt.Sprintf("/%s/", buildPath)
)

func main() {
	handlers := server.NewHandlers(buildPath)

	handlerMux := http.NewServeMux()
	handlerMux.HandleFunc("/api", handlers.HelloWorld)
	handlerMux.HandleFunc("/", handlers.Home)
	handlerMux.Handle(buildURL, http.StripPrefix(buildURL, http.FileServer(http.Dir(buildPath))))

	log.Printf("listening on %s", addr)
	err := http.ListenAndServe(addr, handlerMux)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
}
