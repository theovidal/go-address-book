package cli

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"strconv"
)

// Delete prompts the user to remove a contact using the command line interface
func (instance *Cli) Delete() {
	deletingString, _ := instance.Localizer.Localize(&i18n.LocalizeConfig{
    MessageID: "ContactsDeleting",
	})
	deletedString, _ := instance.Localizer.Localize(&i18n.LocalizeConfig{
    MessageID: "ContactDeleted",
	})

	cancelString, _ := instance.Localizer.Localize(&i18n.LocalizeConfig{
    MessageID: "Cancel",
	})
	cancelingString, _ := instance.Localizer.Localize(&i18n.LocalizeConfig{
    MessageID: "Canceling",
	})

	println(deletingString)
	println(cancelString)
	for _, contact := range instance.Book.ListAllContacts(true) {
		print(contact)
	}
	line, err := instance.Reader.Readline()
	if err != nil {
		panic(err)
	}

	choice, err := strconv.Atoi(line)
	if err != nil {
		panic(err)
	}

	if choice == 0 {
		println(cancelingString)
		return
	}

	instance.Book.DeleteContact(choice)
	println(deletedString)
}
