package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	log.Println("Hello World")

	h := http.NewServeMux()
	h.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})
	h.HandleFunc("/render", func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("template.html")
		content := "I love architecture."
		t.Execute(w, content)
	})
	err := http.ListenAndServe(":8888", h)
	if err != nil {
		log.Fatal(err)
	}
}
