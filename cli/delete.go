package cli

import "strconv"

func (instance *Cli) Delete() {
	println("Indiquez le numéro du contact à supprimer :")
	println("0 : annuler l'opération")
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
		println("Annulation de l'opération")
		return
	}

	instance.Book.DeleteContact(choice)
	println("Contact supprimé")
}
