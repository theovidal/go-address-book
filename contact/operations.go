package contact

import (
	"fmt"
	"github.com/gookit/color"
)

func (contact Contact) String() string {
	var final string

	final += color.Bold.Sprintf("%s <%s>\n", contact.Name, contact.Email)
	final += fmt.Sprintf("- %s\n- %s\n", contact.Address, contact.Phone)

	return final
}