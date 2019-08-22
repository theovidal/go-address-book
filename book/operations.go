package book

import (
	"fmt"
	"os"

	"github.com/chzyer/readline"
	"github.com/exybore/go-address-book/contact"
	"github.com/gocarina/gocsv"
)

func (book Book) Save() {
	file, _ := os.OpenFile("book.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	err := gocsv.MarshalFile(&book.Contacts, file)
	if err != nil {
		panic(err)
	}
}

func (book Book) CreateContact(reader *readline.Instance) {
	promps := map[string]string{
		"Nom": "",
		"Adresse email": "",
		"Adresse physique": "",
		"Numéro de téléphone": "",
	}
	println("Ajout d'un nouveau contact :")
	for field, _ := range promps {
		println(field, ":")
		line, err := reader.Readline()
		if err != nil {
			panic(err)
		}
		promps[field] = line
	}
	book.Contacts = append(book.Contacts, &contact.Contact{
		Name: promps["Nom"],
		Email: promps["Adresse email"],
		Address: promps["Adresse physique"],
		Phone: promps["Numéro de téléphone"],
	})
	println("Contact ajouté")
}

func (book Book) ListAllContacts(withIndex bool) {
	println("Liste des contact :")
	println()
	for index, contact := range book.Contacts {
		fmt.Println(index, ":", contact)
	}
}