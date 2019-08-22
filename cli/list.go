package cli

func (instance Cli) List() {
	println("Liste des contact :")
	for _, contact := range instance.Book.ListAllContacts(true) {
		println(contact)
	}
}
