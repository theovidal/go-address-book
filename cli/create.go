package cli

// Create prompts the user to add a contact using the command line interface
func (instance *Cli) Create() {
	nameString, _ := instance.I18n.T("Name", nil)

	addingString, _ := instance.I18n.T("ContactAdding", nil)
	println(addingString)
	println(nameString, ":")
	line, err := instance.Reader.Readline()
	if err != nil {
		panic(err)
	}

	instance.Book.CreateContact(line)

	addedString, _ := instance.I18n.T("ContactAdded", map[string]interface{}{
		"Name": line,
	})
	println(addedString)
}
