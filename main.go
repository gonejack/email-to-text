package main

import (
	"log"
	"os"

	"github.com/gonejack/email-to-text/cmd"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	err := new(cmd.EmailToText).Run()
	if err != nil {
		log.Fatal(err)
	}
}
