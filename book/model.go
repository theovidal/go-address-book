// Package book provides all the core features of the address book : managing contacts
package book

import (
	"github.com/exybore/go-address-book/contact"
)

// Book structure represents the address book and its data
type Book struct {
	Contacts []*contact.Contact
}
