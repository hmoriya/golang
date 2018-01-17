package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Hello World")

	h := http.NewServeMux()
	h.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})
	err := http.ListenAndServe(":8888", h)
	if err != nil {
		log.Fatal(err)
	}
}
