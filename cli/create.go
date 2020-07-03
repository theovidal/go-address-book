package cli

import "github.com/nicksnyder/go-i18n/v2/i18n"

// Create prompts the user to add a contact using the command line interface
func (instance *Cli) Create() {
	nameString, _ := instance.Localizer.Localize(&i18n.LocalizeConfig{
		MessageID: "Name",
	})

	addingString, _ := instance.Localizer.Localize(&i18n.LocalizeConfig{
		MessageID: "ContactAdding",
	})
	println(addingString)
	println(nameString, ":")
	line, err := instance.Reader.Readline()
	if err != nil {
		panic(err)
	}

	instance.Book.CreateContact(line)

	addedString, _ := instance.Localizer.Localize(&i18n.LocalizeConfig{
		MessageID: "ContactAdded",
		TemplateData: map[string]interface{}{
			"Name": line,
		},
	})
	println(addedString)
}
