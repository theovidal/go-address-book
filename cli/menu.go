package cli

import "github.com/AlecAivazis/survey/v2"

// Menu displays and loops over the menu in the command line interface
func (instance Cli) Menu() {
	openingString, _ := instance.I18n.T("Opening", nil)
	println(openingString)

	actionsString, _ := instance.I18n.T("Actions", nil)
	listString, _ := instance.I18n.T("List", nil)
	createString, _ := instance.I18n.T("Create", nil)
	updateString, _ := instance.I18n.T("Update", nil)
	deleteString, _ := instance.I18n.T("Delete", nil)
	closeString, _ := instance.I18n.T("Close", nil)
	unknownChoiceString, _ := instance.I18n.T("UnknownChoice", nil)

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

	closingString, _ := instance.I18n.T("Closing", nil)
	println(closingString)
	_ = instance.Book.Save()
	_ = instance.Reader.Close()
}
