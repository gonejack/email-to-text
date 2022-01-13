package main

import (
	"log"
	"os"

	"github.com/gonejack/email-to-text/email2text"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	cmd := email2text.EmailToText{
		Options: email2text.MustParseOptions(),
	}
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
