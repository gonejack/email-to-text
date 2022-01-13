package email2text

import (
	"path/filepath"

	"github.com/alecthomas/kong"
)

type Options struct {
	Verbose bool `help:"Verbose printing."`
	About   bool `help:"About."`

	Eml []string `name:".eml" arg:"" optional:"" help:"list of .eml files"`
}

func MustParseOptions() (opt Options) {
	kong.Parse(&opt,
		kong.Name("email-to-text"),
		kong.Description("This command line converts .eml to .txt"),
		kong.UsageOnError(),
	)
	if len(opt.Eml) == 0 || opt.Eml[0] == "*.eml" {
		opt.Eml, _ = filepath.Glob("*.eml")
	}
	return
}
