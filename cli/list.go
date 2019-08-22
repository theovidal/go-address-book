package cli

// Lists the contact into the command line interface
func (instance Cli) List() {
	println("Liste des contact :")
	for _, contact := range instance.Book.ListAllContacts(false) {
		println(contact)
	}
}
