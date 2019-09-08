package cli

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"strconv"
)

// Menu displays and loops over the menu in the command line interface
func (instance Cli) Menu() {
	openingString, _ := instance.Localizer.Localize(&i18n.LocalizeConfig{
		MessageID: "Opening",
	})
	println(openingString)

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

	actions := []string{
		listString,
		createString,
		updateString,
		deleteString,
		closeString,
	}

	quit := false
	for !quit {
		actionsString, _ := instance.Localizer.Localize(&i18n.LocalizeConfig{
			MessageID: "Actions",
		})
		println(actionsString)
		for index, action := range actions {
			println(index+1, ":", action)
		}

		line, err := instance.Reader.Readline()
		if err != nil {
			panic(err)
		}

		chosen, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		switch chosen {
		case 1:
			instance.List()
		case 2:
			instance.Create()
		case 4:
			instance.Delete()
		case 5:
			quit = true
		default:
			unknownChoiceString, _ := instance.Localizer.Localize(&i18n.LocalizeConfig{
				MessageID: "UnknownChoice",
			})
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
