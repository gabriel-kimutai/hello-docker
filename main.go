package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	ticker := time.NewTicker(time.Second * 1)
	go func() {
		for t := range ticker.C {
			log.Printf("current_datetime: %v", t.Format(time.RFC3339))
		}
	}()
	mux.HandleFunc("/hello/{name}", handleHello)
	if err := http.ListenAndServe(":8000", mux); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

}

func handleHello(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	if name == "" {
		writeError(w, http.StatusBadRequest, "Empty name")
		return
	}
	res := fmt.Sprintf("Hello, %s", name)
	writeResponse(w, res)
}

func writeError(w http.ResponseWriter, code int, msg string) {
	w.WriteHeader(code)
	w.Write([]byte(msg))
}

func writeResponse(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(msg))
}
