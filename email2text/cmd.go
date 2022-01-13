package email2text

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gonejack/email"
	"github.com/k3a/html2text"
)

type EmailToText struct {
	Options
}

func (c *EmailToText) Run() (err error) {
	if c.About {
		fmt.Println("Visit https://github.com/gonejack/email-to-text")
		return
	}
	if len(c.Eml) == 0 {
		return fmt.Errorf("no .eml files given")
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

		err = os.WriteFile(target, []byte(content), 0766)
		if err != nil {
			return err
		}
	}

	return nil
}
