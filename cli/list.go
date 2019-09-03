package cli

import "github.com/nicksnyder/go-i18n/v2/i18n"

// Lists the contact into the command line interface
func (instance Cli) List() {
	listingString, _ := instance.Localizer.Localize(&i18n.LocalizeConfig{
    MessageID: "ContactsListing",
	})
	println(listingString)
	for _, contact := range instance.Book.ListAllContacts(false) {
		println(contact)
	}
}
