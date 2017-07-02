package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	filePath := "D:/data/"
	fileServer := http.FileServer(http.Dir(filePath))
	go http.ListenAndServe(":8080", fileServer)

	log.Printf("Listening on: http://localhost:8080")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	log.Println("Shutting down...")
}
