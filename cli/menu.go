package cli

import (
	"strconv"
)

func (instance Cli) Menu() {
	println("Ouverture du carnet d'adresses")

	actions := []string{
		"Lister les contact",
		"Ajouter un contact",
		"Modifier un contact",
		"Supprimer un contact",
		"Fermer le carnet",
	}

	quit := false
	for !quit {
		println("Actions :")
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
		case 5:
			quit = true
		default:
			println("Choix inconnu. Veuillez réessayer.")
		}
	}

	println("Fermeture du carnet d'adresses")
	instance.Book.Save()
}
