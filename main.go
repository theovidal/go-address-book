package main

import (
	"fmt"
	"github.com/exybore/go-address-book/book"
	"github.com/exybore/go-address-book/cli"
	"github.com/exybore/go-address-book/config"
)

func main() {
	configInstance, err := config.NewConfig()
	if err != nil {
		fmt.Println("Error while instancing configuration :", err)
		return
	}

	bookInstance, err := book.New()
	if err != nil {
		fmt.Println("Error while instancing book :", err)
		return
	}

	cliInstance, err := cli.NewInstance(bookInstance, configInstance)
	if err != nil {
		fmt.Println("Error while instancing command-line interface :", err)
		return
	}

	cliInstance.Menu()
}
