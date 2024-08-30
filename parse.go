package harvest

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

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
