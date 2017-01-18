package main

import (
	"html/template"
	"io/ioutil"
)

const layout = "./template/layout.html"

var layoutTemplate *template.Template

func init() {
	layoutTemplate = template.Must(template.New("layout").ParseFiles(layout))
}

func makeTemplate(fname string) (*template.Template, error) {
	tpl := template.Must(layoutTemplate.Clone())
	source, err := ioutil.ReadFile(fname)

	if err != nil {
		return nil, err
	}

	return tpl.New("content").Parse(string(source))
}
