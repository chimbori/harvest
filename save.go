package main

import (
	"encoding/base64"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var (
	includeFilter string
	mimeType      string
	renumber      bool
	minBytes      int
)

func Save(args []string) {
	saveFlags := flag.NewFlagSet("save", flag.ExitOnError)
	saveFlags.StringVar(&includeFilter, "include", "", "only include URLs containing this substring")
	saveFlags.StringVar(&mimeType, "type", "", "only include matching MIME types")
	saveFlags.IntVar(&minBytes, "minbytes", 0, "only include files greater than this")
	saveFlags.BoolVar(&renumber, "renumber", true, "true to renumber; false to keep original filenames")
	saveFlags.Parse(args)

	if len(saveFlags.Args()) == 0 {
		saveFlags.Usage()
		fmt.Fprintf(os.Stderr, "\n%s save [options] <file> <file> …\n\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	firstFilePath, err := filepath.Abs(saveFlags.Args()[0])
	if err != nil {
		log.Fatal(err)
	}
	if _, err := os.Stat(firstFilePath); errors.Is(err, os.ErrNotExist) {
		log.Fatal(err)
	}

	outputLoc := strings.TrimSuffix(firstFilePath, filepath.Ext(firstFilePath))
	filePrefix := filepath.Base(outputLoc)

	os.MkdirAll(outputLoc, os.ModePerm)

	numberSuffix := 1
	for _, fileName := range saveFlags.Args() {
		har, err := Parse(fileName)
		if err != nil {
			log.Println(err)
		}

		for _, entry := range har.Log.Entries {
			if includeFilter != "" && !strings.Contains(entry.Request.URL, includeFilter) {
				log.Println("Skipping", entry.Request.URL)
				continue
			}

			if mimeType != "" && !strings.Contains(entry.Response.Content.MimeType, mimeType) {
				log.Println("Skipping", entry.Request.URL)
				continue
			}

			var content []byte
			if entry.Response.Content.Encoding == "base64" {
				content, _ = base64.StdEncoding.DecodeString(entry.Response.Content.Text)
			} else {
				content = []byte(entry.Response.Content.Text)
			}

			if len(content) < minBytes {
				log.Println("Skipping", entry.Request.URL)
				continue
			}

			fileName, err := getFileName(filePrefix, entry.Request.URL, numberSuffix)
			if err != nil {
				panic(err)
			}
			fileName = filepath.Join(outputLoc, fileName)

			log.Println(entry.Request.URL, " --> ", fileName)
			err = os.WriteFile(fileName, content, 0o644)
			if err != nil {
				log.Println(err)
			}

			numberSuffix++
		}
	}
}

func getFileName(filePrefix string, fileUrl string, numberSuffix int) (fileName string, err error) {
	u, err := url.Parse(fileUrl)
	if err != nil {
		panic(err)
	}

	if renumber {
		ext := path.Ext(u.Path)
		if ext == "" {
			ext = ".html"
		}
		return fmt.Sprintf("%s %d%s", filePrefix, numberSuffix, ext), nil
	} else {
		return path.Base(u.Path), nil
	}
}
