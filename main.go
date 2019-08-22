package main

import (
	"os"

	"github.com/exybore/go-address-book/book"
	"github.com/exybore/go-address-book/cli"
	"github.com/exybore/go-address-book/contact"
	"github.com/gocarina/gocsv"
)

func main() {
	file, err := os.OpenFile("book.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var contactsList []*contact.Contact

	if err := gocsv.UnmarshalFile(file, &contactsList); err != nil {
		panic(err)
	}

	bookInstance := book.Book{Contacts: contactsList}

	cliInstance, err := cli.NewInstance(bookInstance)
	if err != nil {
		panic(err)
	}

	cliInstance.Menu()

	cliInstance.Reader.Close()
}
