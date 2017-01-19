package main

import (
	"bytes"
	"html/template"
	"io"
	"log"
	"net/http"
)

type downloadsHandler struct {
	Downloads *Downloads
	template  *template.Template
}

func enqueueDownload(d *Downloads, w http.ResponseWriter, r *http.Request) {
	src := r.FormValue("src")

	dst, err := pickFilename(src)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Enqueued %s to %s", src, dst)
	d.enqueue(src, dst)

	http.Redirect(w, r, "/", http.StatusFound)
}

func renderTemplate(t *template.Template, w http.ResponseWriter, name string, data interface{}) {
	buf := bytes.NewBuffer(make([]byte, 0))
	err := t.ExecuteTemplate(buf, name, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Add("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		io.Copy(w, buf)
	}
}

func (d downloadsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		renderTemplate(d.template, w, "list.html", d.Downloads)
	case http.MethodPost:
		enqueueDownload(d.Downloads, w, r)
	}
}
