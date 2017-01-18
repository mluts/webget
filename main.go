package main

import (
	"html/template"
	"log"
	"net/http"
)

const layout = `
<html>
	<head></head>
	<body>
		{{template "content"}}
	</body>
</html>
`

const index = `
<h1>
	{{.}}
</h1>
`

const addr = ":8080"

func main() {
	layout := template.Must(template.New("layout").Parse(layout))

	tpl := template.Must(layout.Clone())
	template.Must(tpl.Parse(index))

	data := "hello world"

	handler := MethodHandler{
		http.MethodGet: TemplateHandler{
			tpl, &data}}

	http.Handle("/", handler)
	log.Printf("Initializing server at %s", addr)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
