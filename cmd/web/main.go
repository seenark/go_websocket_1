package main

import (
	"log"
	"net/http"
	"ws/internal/handlers"
)

func main() {
	mux := routes()
	go handlers.ListenToWsChan()
	log.Println("Staring web server on port 8080")

	_ = http.ListenAndServe(":8080", mux)
}
