// Package cli contains all the command line interactions features
package cli

import (
	"github.com/BurntSushi/toml"
	"github.com/chzyer/readline"
	"github.com/exybore/go-address-book/book"
	"github.com/exybore/go-address-book/config"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

// Cli structure represents the command line interface
type Cli struct {
	Book      book.Book
	Reader    *readline.Instance
	Config    config.Config
	Localizer *i18n.Localizer
}

// NewInstance returns an instance of the Cli structure
func NewInstance(book book.Book, conf config.Config) (cli Cli, err error) {
	reader, err := readline.New("> ")
	if err != nil {
		return
	}

	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	if _, err = bundle.LoadMessageFile("i18n/cli/en.toml"); err != nil {
		return
	}
	if _, err = bundle.LoadMessageFile("i18n/cli/fr.toml"); err != nil {
		return
	}

	localizer := i18n.NewLocalizer(bundle, conf.Locale)

	cli = Cli{
		Book:      book,
		Reader:    reader,
		Config:    conf,
		Localizer: localizer,
	}
	return
}
