package file

import (
	"flag"
	"log"
	"os"
)

func Read() []byte {
	file := flag.Args()
	if len(file) == 0 {
		log.Fatal("No file input")
	}

	text, err := os.ReadFile(file[0])
	if err != nil {
		log.Fatal(err)
	}

	return text
}
