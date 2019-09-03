// The cli package contains all the command line interactions features
package cli

import (
	"github.com/BurntSushi/toml"
	"github.com/chzyer/readline"
	"github.com/exybore/go-address-book/book"
	"github.com/exybore/go-address-book/config"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

// The structure for the command line interface
type Cli struct {
	Book        book.Book
	Reader      *readline.Instance
	Config      config.Config
	Localizer   *i18n.Localizer
}

// Returns an instance of the Cli structure
func NewInstance(book book.Book, conf config.Config) (Cli, error) {
	reader, err := readline.New("> ")
	if err != nil {
		return Cli{}, err
	}

	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	if _, err = bundle.LoadMessageFile("i18n/cli/en.toml"); err != nil {
		return Cli{}, err
	}
	if _, err = bundle.LoadMessageFile("i18n/cli/fr.toml"); err != nil {
		return Cli{}, err
	}

	localizer := i18n.NewLocalizer(bundle, conf.Locale)

	return Cli{
		Book: book,
		Reader: reader,
		Config: conf,
		Localizer: localizer,
	}, nil
}
