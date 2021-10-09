package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/alecthomas/kong"
	"github.com/gonejack/email"
	"github.com/k3a/html2text"
)

type options struct {
	Verbose bool `help:"Verbose printing."`

	Eml []string `arg:"" optional:""`
}

type EmailToText struct {
	options
}

func (c *EmailToText) Run() (err error) {
	kong.Parse(&c.options,
		kong.Name("email-to-text"),
		kong.Description("Convert .eml to .txt"),
		kong.UsageOnError(),
	)
	if len(c.Eml) == 0 {
		c.Eml, _ = filepath.Glob("*.eml")
	}
	if len(c.Eml) == 0 {
		return fmt.Errorf("no email given")
	}

	return c.run()
}

func (c *EmailToText) run() error {
	for _, eml := range c.Eml {
		if c.Verbose {
			log.Printf("processing %s", eml)
		}

		data, err := os.ReadFile(eml)
		if err != nil {
			return err
		}

		mail, err := email.NewEmailFromReader(bytes.NewReader(data))
		if err != nil {
			return err
		}

		target := strings.TrimSuffix(eml, filepath.Ext(eml)) + ".txt"
		content := string(mail.Text)
		if len(content) == 0 {
			content = html2text.HTML2Text(string(mail.HTML))
		}

		err = ioutil.WriteFile(target, []byte(content), 0666)
		if err != nil {
			return err
		}
	}

	return nil
}
