package cli

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

// Menu displays and loops over the menu in the command line interface
func (instance Cli) Menu() {
	openingString, _ := instance.Localizer.Localize(&i18n.LocalizeConfig{
		MessageID: "Opening",
	})
	println(openingString)

	actionsString, _ := instance.Localizer.Localize(&i18n.LocalizeConfig{
		MessageID: "Actions",
	})
	listString, _ := instance.Localizer.Localize(&i18n.LocalizeConfig{
		MessageID: "List",
	})
	createString, _ := instance.Localizer.Localize(&i18n.LocalizeConfig{
		MessageID: "Create",
	})
	updateString, _ := instance.Localizer.Localize(&i18n.LocalizeConfig{
		MessageID: "Update",
	})
	deleteString, _ := instance.Localizer.Localize(&i18n.LocalizeConfig{
		MessageID: "Delete",
	})
	closeString, _ := instance.Localizer.Localize(&i18n.LocalizeConfig{
		MessageID: "Close",
	})
	unknownChoiceString, _ := instance.Localizer.Localize(&i18n.LocalizeConfig{
		MessageID: "UnknownChoice",
	})

	var choice int
	prompt := &survey.Select{
		Message: actionsString,
		Options: []string{
			listString,
			createString,
			updateString,
			deleteString,
			closeString,
		},
	}

	quit := false
	for !quit {
		_ = survey.AskOne(prompt, &choice)

		switch choice {
		case 0:
			instance.List()
		case 1:
			instance.Create()
		case 3:
			instance.Delete()
		case 4:
			quit = true
		default:
			println(unknownChoiceString)
		}
	}

	closingString, _ := instance.Localizer.Localize(&i18n.LocalizeConfig{
		MessageID: "Closing",
	})
	println(closingString)
	instance.Book.Save()
	instance.Reader.Close()
}
