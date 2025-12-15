package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hi from the Go container!")
	})
	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
