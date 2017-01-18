package main

import (
	"html/template"
	"log"
	"net/http"
)

const addr = ":8080"

func main() {
	data := "hello world"

	handler := MethodHandler{
		http.MethodGet: TemplateHandler{
			template.Must(makeTemplate("./template/index.html")), &data}}

	http.Handle("/", handler)
	log.Printf("Initializing server at %s", addr)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
