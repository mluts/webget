package main

import (
	"bufio"
	"flag"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

const defaultAddr = ":8080"
const defaultDownloadDir = "."

var (
	addr        string
	downloadDir string
	help        bool
)

func init() {
	flag.BoolVar(&help, "h", false, "See this message")
	flag.StringVar(&addr, "a", defaultAddr, "Listen address")
	flag.StringVar(&downloadDir, "d", downloadDir, "Download dir")
}

func downloadHTTP(dst, src string) error {
	var err error

	err = os.MkdirAll(downloadDir, 0700)
	if err != nil {
		return err
	}

	dst = path.Join(downloadDir, dst)

	file, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0600)
	if err != nil {
		return err
	}

	res, err := http.Get(src)
	if err != nil {
		return err
	}

	out := bufio.NewWriter(file)
	_, err = io.Copy(out, res.Body)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	flag.Parse()

	if help {
		flag.Usage()
		os.Exit(2)
	}

	http.Handle("/", downloadsHandler{
		&Downloads{GetFile: downloadHTTP},
		template.Must(template.New("list.html").Parse(listTemplate))})

	log.Printf("Listening at %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

const listTemplate = `
{{define "download"}}
	<p>
		{{.Url}}
	</p>
	<p>
		Status:
		{{if .Status | eq 0}}
			Enqueued
		{{else if .Status  | eq 1}}
			Started
		{{else if .Status | eq 2}}
			Success
		{{end}}
	</p>
{{end}}

<html>
	<head>
		<title>Downloads</title>
	</head>

	<body>
		<h2>Enqueue</h2>
		<form method="post" action="/">
			<input type="text" name="src">
			<input type="submit" value="Send">
		</form>

		<h2>Downloads:</h2>
		<ul>
			{{range .List}}
				<li>
					{{template "download" .}}
				</li>
			{{end}}
		</ul>
		</div>
	</body>
</html>
`
