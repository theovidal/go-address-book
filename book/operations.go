package book

import (
	"fmt"
	"os"

	"github.com/exybore/go-address-book/contact"
	"github.com/gocarina/gocsv"
)

// Saves the address book into the file
func (book Book) Save() {
	err := os.Truncate("book.csv", 0)
	if err != nil {
		panic(err)
	}

	file, _ := os.OpenFile("book.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	err = gocsv.MarshalFile(&book.Contacts, file)
	if err != nil {
		panic(err)
	}
}

// Creates a contact and adds it into the book
func (book *Book) CreateContact(params map[string]string) {
	book.Contacts = append(book.Contacts, &contact.Contact{
		Name: params["Name"],
		Email: params["Email"],
		Address: params["Address"],
		Phone: params["Phone"],
	})
}

// Deletes a contact from the book
func (book *Book) DeleteContact(index int) {
	index -= 1
	book.Contacts = append(book.Contacts[:index], book.Contacts[index+1:]...)
}

// Lists all the contacts in a pretty way
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