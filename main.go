package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("Starting Hello World")

	log.Fatal(http.ListenAndServe(hostPort(), http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})))
}

func hostPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return fmt.Sprintf(":%s", port)
}
