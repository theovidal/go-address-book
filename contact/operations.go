package contact

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gookit/color"
)

// Strings displays the contact as a readable string
func (contact Contact) String() string {
	var final string

	final += color.Bold.Sprintf("%s <%s>\n", contact.Name, contact.Emails[0].Value)
	// final += fmt.Sprintf("- %s : %s\n", contact.Address, contact.Phone)
	final += fmt.Sprintf("- %s\n", contact.Phones)

	return final
}

func GenerateUUID() string {
	return uuid.New().String()
}
