package main

import (
	"github.com/coderman-engineering/hello-world/pkg/greet"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", defaultGreeting)
	http.ListenAndServe("0.0.0.0:4000", mux)
}

func defaultGreeting(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(greet.DefaultGreeting()))
	w.WriteHeader(http.StatusOK)
}
