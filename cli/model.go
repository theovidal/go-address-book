// Package cli contains all the command line interactions features
package cli

import (
	"github.com/chzyer/readline"
	"github.com/exybore/go-address-book/book"
	"github.com/exybore/go-address-book/config"
	"github.com/exybore/go-address-book/i18n"
)

// Cli structure represents the command line interface
type Cli struct {
	Book   book.Book
	Reader *readline.Instance
	Config config.Config
	I18n   *i18n.I18n
}

// NewInstance returns an instance of the Cli structure
func NewInstance(book book.Book, conf config.Config) (cli Cli, err error) {
	reader, err := readline.New("> ")
	if err != nil {
		return
	}

	i18nInstance, err := i18n.NewInstance(conf.Locale)

	cli = Cli{
		Book:   book,
		Reader: reader,
		Config: conf,
		I18n:   &i18nInstance,
	}
	return
}
