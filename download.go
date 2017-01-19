package main

import (
	"github.com/hashicorp/go-getter"
	"log"
)

const (
	Enqueued = 0
	Started  = 1
	Success  = 2
	Failure  = 3
)

type Download struct {
	Url, Fname string
	Status     int
	Err        error
}

type Downloads struct {
	List []Download
}

func downloadTo(url string, fname string) error {
	return getter.GetFile(fname, url)
}

func (d *Downloads) enqueue(url string, fname string) {
	d.List = append(d.List, Download{
		Url:    url,
		Fname:  fname,
		Status: Enqueued})
	go d.download(&d.List[len(d.List)-1])
}

func (d *Downloads) download(t *Download) {
	log.Printf("Starting to download %v", t)
	(*t).Status = Started

	err := downloadTo(t.Url, t.Fname)
	if err != nil {
		log.Printf("Failed: %v", err)
		t.Status = Failure
		t.Err = err
	} else {
		log.Printf("Success!")
		t.Status = Success
	}
}
