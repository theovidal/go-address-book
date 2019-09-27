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
	Emails map[string]string
	// One or more physical addresses
	Addresses map[string]Address
	// One or more phone numbers
	Phones map[string]string
	// One or more Internet links
	WebLinks []string
	// One or more accounts on Internet
	Accounts map[string]string
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

type Event struct {
	// The type of the event : birthday...
	Type string
	// The time when this event happens
	Time time.Time
}
