package main

import (
	"github.com/exybore/go-address-book/book"
	"github.com/exybore/go-address-book/cli"
	"github.com/exybore/go-address-book/config"
)

func main() {
	configInstance, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	bookInstance := book.New()

	cliInstance, err := cli.NewInstance(bookInstance, configInstance)
	if err != nil {
		panic(err)
	}

	cliInstance.Menu()
}
