package main

import (
	"fmt"
	"net/http"
)

const port = ":3001"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from your first http server")
	})
	fmt.Printf("Starting server in port %s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}
