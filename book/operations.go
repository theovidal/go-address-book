package book

import (
	"fmt"
	"os"

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

func (book *Book) CreateContact(params map[string]string) {
	book.Contacts = append(book.Contacts, &contact.Contact{
		Name: params["Name"],
		Email: params["Email"],
		Address: params["Address"],
		Phone: params["Phone"],
	})
}

func (book Book) ListAllContacts(withIndex bool) []string {
	var contacts []string
	for index, contact := range book.Contacts {
		var prefix string
		if withIndex {
			prefix += fmt.Sprintf("%d : ", index + 1)
		}
		contacts = append(contacts, fmt.Sprint(prefix, contact))
	}
	return contacts
}