package main

import (
	"html/template"
	"net/http"
)

type MethodHandler map[string]http.Handler

func (m MethodHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	handler, ok := m[req.Method]
	if !ok {
		w.WriteHeader(405)
		return
	}

	handler.ServeHTTP(w, req)
}

type TemplateHandler struct {
	template *template.Template
	data     interface{}
}

func (h TemplateHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h.template.Execute(w, h.data)
}
