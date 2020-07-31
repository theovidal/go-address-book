package cli

import "github.com/AlecAivazis/survey/v2"

// Delete prompts the user to remove a contact using the command line interface
func (instance *Cli) Delete() {
	deletingString, _ := instance.I18n.T("ContactDeleting", nil)

	cancelString, _ := instance.I18n.T("Cancel", nil)

	choices := []string{cancelString}

	for _, contact := range instance.Book.Contacts {
		choices = append(choices, contact.Name)
	}

	var choice int
	prompt := &survey.Select{
		Message: deletingString,
		Options: choices,
	}
	_ = survey.AskOne(prompt, &choice)

	if choice == 0 {
		return
	}

	deletedString, _ := instance.I18n.T("ContactDeleted", map[string]interface{}{
		"Name": choices[choice],
	})

	instance.Book.DeleteContact(choice)
	println(deletedString)
}
