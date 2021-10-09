package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Word")
}

func main() {
	server := &http.Server{
		Addr: "149.28.124.38:8999",
	}
	http.HandleFunc("/", hello)
	server.ListenAndServe()
}
