// The cli package contains all the command line interactions features
package cli

import (
	"github.com/chzyer/readline"
	"github.com/exybore/go-address-book/book"
)

// The structure for the command line interface
type Cli struct {
	Book   book.Book
	Reader *readline.Instance
}

// Returns an instance of the Cli structure
func NewInstance(book book.Book) (Cli, error) {
	reader, err := readline.New("> ")
	if err != nil {
		return Cli{}, err
	}

	return Cli{Book: book, Reader: reader}, nil
}
