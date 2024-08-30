package save

import (
	"encoding/base64"
	"log"
	"net/url"
	"os"
	"path"

	"go.chimbori.app/harvest"
)

func Save(args []string) {
	log.Printf("Save")

	for _, fileName := range args {
		har, err := harvest.Parse(fileName)
		if err != nil {
			log.Println(err)
		}

		for _, entry := range har.Log.Entries {
			// log.Println(entry.Response.Content.MimeType)
			// log.Println(entry.Request.URL)
			// log.Println(entry.Response.Content.Text)

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

			log.Println(entry.Request.URL, " --> ", fileName)
			err = os.WriteFile("tmp-"+fileName, content, 0o644)
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
