package main

import (
	"bufio"
	"io"
	"net/http"
	"os"
)

func downloadTo(url string, fname string) (int64, error) {
	file, err := os.OpenFile(fname, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
	defer file.Close()
	if err != nil {
		return 0, err
	}

	res, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	return io.Copy(bufio.NewWriter(file), res.Body)
}
