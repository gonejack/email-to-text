package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gonejack/email"
	"github.com/k3a/html2text"
	"github.com/mvdan/xurls"
	"github.com/spf13/cobra"
)

var (
	pattern, _ = xurls.StrictMatchingScheme("https?://")
	verbose    = false
	prog       = &cobra.Command{
		Use:   "email-links *.eml",
		Short: "Command line tool for extract links from emails",
		Run: func(c *cobra.Command, args []string) {
			err := run(c, args)
			if err != nil {
				log.Fatal(err)
			}
		},
	}
)

func init() {
	log.SetOutput(os.Stdout)

	prog.Flags().SortFlags = false
	prog.PersistentFlags().SortFlags = false
	prog.PersistentFlags().BoolVarP(
		&verbose,
		"verbose",
		"v",
		false,
		"verbose",
	)
}

func run(c *cobra.Command, emails []string) error {
	if len(emails) == 0 {
		emails, _ = filepath.Glob("*.eml")
	}
	if len(emails) == 0 {
		return fmt.Errorf("no email given")
	}

	for _, eml := range emails {
		if verbose {
			log.Printf("processing %s", eml)
		}

		fd, err := os.Open(eml)
		if err != nil {
			return err
		}

		mail, err := email.NewEmailFromReader(fd)
		if err != nil {
			return err
		}

		_ = fd.Close()

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

func printURLs(text string) {
	if text == "" {
		return
	}
	for _, url := range pattern.FindAllString(text, -1) {
		_, _ = fmt.Fprintln(os.Stdout, url)
	}
}

func main() {
	_ = prog.Execute()
}
