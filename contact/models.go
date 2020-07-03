// Package contact contains all structures and tools related to contacts
package contact

import "time"

// Contacts represents a contact and all its data in the address book
type Contact struct {
	// The Unique User identifier of the user (based on UUIDv4)
	UUID string
	// The full name of the contact
	Name string
	// One or more email addresses
	Emails []TypeValue
	// One or more physical addresses
	Addresses []Address
	// One or more phone numbers
	Phones []TypeValue
	// One or more Internet links
	WebLinks []string
	// One or more accounts on Internet
	Accounts []TypeValue
	// One or more associated contacts
	AssociatedContacts []int
	// Some texts that user can add to the contact
	Notes []string
}

// Address represents a physical address for a contact
type Address struct {
	// The type of the address : personal, professional...
	Type string

	// Next fields are precisions for the address
	Street  string
	City    string
	State   string
	Zip     string
	Country string
}

type TypeValue struct {
	// The type of the field : personal, professional...
	Type string
	// The value of this piece of information
	Value string
}

type Event struct {
	// The type of the event : birthday...
	Type string
	// The time when this event happens
	Time time.Time
}
