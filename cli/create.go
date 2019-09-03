package cli

import "github.com/nicksnyder/go-i18n/v2/i18n"

// Creates a contact using the command line interface
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

	promps := map[string]string{
		nameString: "",
		emailString: "",
		addressString: "",
		phoneString: "",
	}

	addingString, _ := instance.Localizer.Localize(&i18n.LocalizeConfig{
    MessageID: "ContactAdding",
	})
	println(addingString)
	for field, _ := range promps {
		println(field, ":")
		line, err := instance.Reader.Readline()
		if err != nil {
			panic(err)
		}
		promps[field] = line
	}

	params := map[string]string{
		"Name": promps[nameString],
		"Email": promps[emailString],
		"Address": promps[addressString],
		"Phone": promps[phoneString],
	}
	instance.Book.CreateContact(params)

	addedString, _ := instance.Localizer.Localize(&i18n.LocalizeConfig{
    MessageID: "ContactAdded",
    TemplateData: map[string]interface{}{
    	"Name": promps[nameString],
    },
	})
	println(addedString)
}
