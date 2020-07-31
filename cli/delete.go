package cli

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

// Delete prompts the user to remove a contact using the command line interface
func (instance *Cli) Delete() {
	deletingString, _ := instance.Localizer.Localize(&i18n.LocalizeConfig{
		MessageID: "ContactDeleting",
	})

	cancelString, _ := instance.Localizer.Localize(&i18n.LocalizeConfig{
		MessageID: "Cancel",
	})

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

	deletedString, _ := instance.Localizer.Localize(&i18n.LocalizeConfig{
		MessageID: "ContactDeleted",
		TemplateData: map[string]interface{}{
			"Name": choices[choice],
		},
	})

	instance.Book.DeleteContact(choice)
	println(deletedString)
}
