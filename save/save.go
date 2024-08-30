package save

import (
	"encoding/base64"
	"flag"
	"log"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"

	"go.chimbori.app/harvest"
)

var (
	outputLoc     string
	includeFilter string
	mimeType      string
)

func Save(args []string) {
	log.Printf("Save")

	saveFlags := flag.NewFlagSet("save", flag.ExitOnError)
	saveFlags.StringVar(&outputLoc, "output", "./", "location to save files to")
	saveFlags.StringVar(&includeFilter, "include", "", "only include URLs containing this substring")
	saveFlags.StringVar(&mimeType, "type", "", "only include matching MIME types")
	saveFlags.Parse(args)

	os.MkdirAll(outputLoc, os.ModePerm)

	for _, fileName := range saveFlags.Args() {
		har, err := harvest.Parse(fileName)
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

			fileName, err := getFileName(entry.Request.URL)
			if err != nil {
				log.Println(err)
			}
			fileName = filepath.Join(outputLoc, fileName)

			log.Println(entry.Request.URL, " --> ", fileName)
			err = os.WriteFile(fileName, content, 0o644)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func getFileName(s string) (fileName string, err error) {
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	fileName = path.Base(u.Path)
	return
}
