package main

import (
	"html/template"
	"log"
	"net/http"
)

const addr = ":8080"

func main() {
	http.Handle("/", downloadsHandler{
		&Downloads{},
		template.Must(template.New("").ParseGlob("./template/*.html"))})

	log.Printf("Listening at %s", addr)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
