package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var commands = map[string]func([]string){
	"save": Save,
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, usage())
		os.Exit(1)
	}

	cmd, ok := commands[os.Args[1]]
	if !ok {
		fmt.Fprintln(os.Stderr, usage())
		os.Exit(1)
	}

	cmd(os.Args[2:])
}

func usage() string {
	s := fmt.Sprintf("Usage: %s <command> [options]\nAvailable commands:\n", filepath.Base(os.Args[0]))
	for k := range commands {
		s += " - " + k + "\n"
	}
	return s
}

func Parse(harFile string) (har Har, err error) {
	file, err := os.Open(harFile)
	if err != nil {
		log.Println(err)
	}

	buf := bufio.NewReader(file)
	dec := json.NewDecoder(buf)
	err = dec.Decode(&har)
	return
}

type Har struct {
	Log struct {
		Entries []struct {
			ResourceType string `json:"_resourceType"`
			Request      struct {
				URL string `json:"url"`
			} `json:"request"`
			Response struct {
				Content struct {
					Encoding string `json:"encoding"`
					MimeType string `json:"mimeType"`
					Size     int64  `json:"size"`
					Text     string `json:"text"`
				} `json:"content"`
			} `json:"response"`
		} `json:"entries"`
	} `json:"log"`
}
