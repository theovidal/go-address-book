package cli

import (
	"github.com/chzyer/readline"
	"github.com/exybore/go-address-book/book"
)

type Cli struct {
	Book   book.Book
	Reader *readline.Instance
}

func NewInstance(book book.Book) (Cli, error) {
	reader, err := readline.New("> ")
	if err != nil {
		return Cli{}, err
	}

	return Cli{Book: book, Reader: reader}, nil
}
