package main

import (
	"github.com/exybore/go-address-book/book"
	"github.com/exybore/go-address-book/cli"
)

func main() {
	bookInstance := book.New()

	cliInstance, err := cli.NewInstance(bookInstance)
	if err != nil {
		panic(err)
	}

	cliInstance.Menu()
}
