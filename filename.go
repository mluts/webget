package main

import (
	"log"
	"net/url"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
)

func fileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func incrementNumber(basePath string) string {
	r, err := regexp.Compile("\\d+$")
	if err != nil {
		panic(err)
	}

	if r.MatchString(basePath) {
		number, err := strconv.ParseInt(r.FindString(basePath), 10, 32)
		if err != nil {
			number = 999
		} else {
			number++
		}

		basePath = r.ReplaceAllString(basePath, strconv.Itoa(int(number)))
	} else {
		basePath = strings.Join([]string{basePath, "0"}, "")
	}

	return basePath
}

func pickUnusedFilename(basePath string) string {
	path := basePath
	for fileExists(path) {
		log.Printf("%s already exists", path)
		path = incrementNumber(path)
		log.Printf("Picked %s", path)
	}
	return path
}

func pickFilename(uri string) (string, error) {
	parsed, err := url.Parse(uri)
	fname := parsed.Path
	fname = strings.TrimSpace(fname)
	fname = strings.TrimPrefix(fname, "/")
	fname = path.Base(fname)

	if err != nil {
		return "", err
	}

	if fname == "." || fname == "/" || fname == "" {
		fname = "unnamed"
	}

	fname = pickUnusedFilename(fname)

	return fname, nil
}
