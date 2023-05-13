package main

import (
	"log"
	"net/http"
	"time"
)

func api(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	time.Sleep(10 * time.Second) // Time consuming task
	w.WriteHeader(http.StatusAccepted)
	ctx.Done()
}

func main() {
	m := http.NewServeMux()
	m.HandleFunc("/api/v1", api)
	err := http.ListenAndServe(":9000", m)
	log.Fatal(err)
}
