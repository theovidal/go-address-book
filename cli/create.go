package cli

import "github.com/nicksnyder/go-i18n/v2/i18n"

// Create prompts the user to add a contact using the command line interface
func (instance *Cli) Create() {
	nameString, _ := instance.Localizer.Localize(&i18n.LocalizeConfig{
		MessageID: "Name",
	})
	emailString, _ := instance.Localizer.Localize(&i18n.LocalizeConfig{
		MessageID: "Email",
	})
	addressString, _ := instance.Localizer.Localize(&i18n.LocalizeConfig{
		MessageID: "Address",
	})
	phoneString, _ := instance.Localizer.Localize(&i18n.LocalizeConfig{
		MessageID: "Phone",
	})

	prompts := map[string]string{
		nameString:    "",
		emailString:   "",
		addressString: "",
		phoneString:   "",
	}

	addingString, _ := instance.Localizer.Localize(&i18n.LocalizeConfig{
		MessageID: "ContactAdding",
	})
	println(addingString)
	for field := range prompts {
		println(field, ":")
		line, err := instance.Reader.Readline()
		if err != nil {
			panic(err)
		}
		prompts[field] = line
	}

	params := map[string]string{
		"Name":    prompts[nameString],
		"Email":   prompts[emailString],
		"Address": prompts[addressString],
		"Phone":   prompts[phoneString],
	}
	instance.Book.CreateContact(params)

	addedString, _ := instance.Localizer.Localize(&i18n.LocalizeConfig{
		MessageID: "ContactAdded",
		TemplateData: map[string]interface{}{
			"Name": prompts[nameString],
		},
	})
	println(addedString)
}
