package book

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/exybore/go-address-book/contact"
	"gopkg.in/yaml.v2"
)

// New initializes a Book structure with all the contacts
func New() (book Book, err error) {
	data, err := ioutil.ReadFile("book.yml")
	if err != nil {
		return
	}

	var contacts []*contact.Contact

	if err = yaml.Unmarshal(data, &contacts); err != nil {
		return
	}

	book = Book{Contacts: contacts}
	return
}

// Save writes the address book into the appropriate file
func (book Book) Save() (err error) {
	err = os.Truncate("book.yml", 0)
	if err != nil {
		return
	}

	file, _ := os.OpenFile("book.yml", os.O_RDWR|os.O_CREATE, os.ModePerm)
	data, err := yaml.Marshal(&book.Contacts)
	if err != nil {
		return
	}

	_, err = file.Write(data)
	if err != nil {
		return
	}

	err = file.Close()
	return
}

// CreateContact add a new contact into the book
func (book *Book) CreateContact(params map[string]string) {
	book.Contacts = append(book.Contacts, &contact.Contact{
		Name:    params["Name"],
		Email:   params["Email"],
		Address: params["Address"],
		Phone:   params["Phone"],
	})
}

// DeleteContact removes a contact from the book
func (book *Book) DeleteContact(index int) {
	index -= 1
	book.Contacts = append(book.Contacts[:index], book.Contacts[index+1:]...)
}

// ListAllContacts returns a list of all the contacts in a pretty way
func (book Book) ListAllContacts(withIndex bool) []string {
	var contacts []string
	for index, actualContact := range book.Contacts {
		var prefix string
		if withIndex {
			prefix += fmt.Sprintf("%d : ", index+1)
		}
		contacts = append(contacts, fmt.Sprint(prefix, actualContact))
	}
	return contacts
}
