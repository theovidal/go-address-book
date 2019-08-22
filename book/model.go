// The package book provides all the core features of the address book : managing contacts
package book

import (
	"github.com/exybore/go-address-book/contact"
)

// The structure for the address book
type Book struct {
	Contacts []*contact.Contact
}
