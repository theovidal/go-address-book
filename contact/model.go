// The contact package contains contacts' structure and tools
package contact

// The structure for a contact in the address book
type Contact struct {
	Name string `csv:"name"`
	Email string `csv:"email"`
	Address string `csv:"address"`
	Phone string `csv:"phone"`
}
