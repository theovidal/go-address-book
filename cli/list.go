package cli

// List displays all the contacts using the command line interface
func (instance Cli) List() {
	listingString, _ := instance.I18n.T("ContactsListing", nil)
	println(listingString)
	for _, contact := range instance.Book.ListAllContacts(false) {
		println(contact)
	}
}
