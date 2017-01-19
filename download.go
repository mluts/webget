package main

import (
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
	List    []Download
	GetFile FileGetter
}

type FileGetter func(dst, src string) error

func (d *Downloads) enqueue(url string, fname string) {
	d.List = append(d.List, Download{
		Url:    url,
		Fname:  fname,
		Status: Enqueued})
	go d.download(&d.List[len(d.List)-1])
}

func (d *Downloads) download(t *Download) {
	log.Printf("Starting to download %s to %s", t.Url, t.Fname)
	t.Status = Started

	if d.GetFile == nil {
		panic("Getter is not set")
	}

	err := d.GetFile(t.Fname, t.Url)
	if err != nil {
		log.Printf("Failed: %v", err)
		t.Status = Failure
		t.Err = err
	} else {
		log.Printf("Success!")
		t.Status = Success
	}
}
