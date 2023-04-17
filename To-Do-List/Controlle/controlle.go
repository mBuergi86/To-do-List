package Controlle

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"To-Do-List/Model"
	"To-Do-List/View"
)

// SignIn → Login
func SignIn() {
	View.Clear()
	count := 3
	fmt.Println("######Benutzername und Passwort abfragen######")
	fmt.Println("")
	for count > 0 {
		fmt.Print("Benutzername: ")
		user := View.AskForInput()
		fmt.Print("Passwort: ")
		password := View.PasswordOfInput()

		if Model.CheckUser(user, password) {
			ViewMenu()
		} else {
			count--
			fmt.Printf("Der Benutzername oder das Passwort sind ungültig. Bitte überprüfen Sie die Eingabe und haben Sie übrig %d Versuche unternommen.\n", count)
		}
	}

	os.Exit(0)
}

// SignUp → Registrierung
func SignUp() {
	View.Clear()
	fmt.Println("#######Bitte registrieren Sie sich neu########")
	fmt.Println("")
	fmt.Print("Benutzername: ")
	user := View.AskForInput()
	fmt.Print("Passwort: ")
	password := View.PasswordOfInput()

	signUp := []Model.Account{
		{
			Username: user,
			Password: password,
		},
	}

	Model.RegisterUser(&signUp)
	View.AccountMenu()
}

// ViewAccount → Die Zugangssicherheit in dem Menü
func ViewAccount() {
	View.AccountMenu()

	var check int

	for {
		input := View.AskForInput()

		if View.IsNumeric(input) {
			check = StringToInteger(input)
		} else {
			fmt.Println("Is Numeric ist false! Bitte geben Sie eine gültige Eingabe ein.")
			continue
		}

		if check < 1 || check > 3 {
			fmt.Println("Bitte geben Sie eine gültige Eingabe ein.")
			continue
		} else {
			SwitchAccount(check)
		}
	}
}

// SwitchAccount → Steuerung von Zugangssicherheit
func SwitchAccount(input int) {
	switch input {
	case 1:
		SignIn()
	case 2:
		SignUp()
	case 3:
		fmt.Println("Die To-do-Liste wird beendet.")
		os.Exit(0)
	}
}

// ViewMenu → Bedienungsmenü
func ViewMenu() {
	View.Menu()

	var check int

	for {
		input := View.AskForInput()

		if View.IsNumeric(input) {
			check = StringToInteger(input)
		} else {
			fmt.Println("Bitte geben Sie eine gültige Eingabe ein.")
			continue
		}

		if check < 1 || check > 5 {
			fmt.Println("Bitte geben Sie eine gültige Eingabe ein.")
			continue
		} else {
			SwitchChoose(check)
		}
	}
}

// SwitchChoose → Steuerung von Bedienungsmenü
func SwitchChoose(input int) {
	switch input {
	case 1:
		AddTask()
	case 2:
		EditTask()
	case 3:
		RemoveTask()
	case 4:
		LoadTaskAll()
	case 5:
		fmt.Println("Die To-do-Liste wird beendet.")
		os.Exit(0)
	}
}

// AddTask → Die Aufgabe hinzufügen
func AddTask() {
	View.Clear()

	var (
		input string
		tasks []Model.Task
	)

	fmt.Println("############################################################################")

	Model.AllList()

	fmt.Println("############################################################################")

	for {
		var (
			status string
		)

		fmt.Println("Bitte geben Sie die neue Aufgabe ein oder geben Sie 'q' ein, um zu beenden:")

		input = View.AskForInput()

		if input == "q" {
			break
		}

		fmt.Print("Klassifizierung:\n[1] Backlog\n[2] To-Do\n[3] Ongoing\n[4] Done\nBitte wählen Sie die Status: ")

		var (
			n int
		)

		for {
			number := View.AskForInput()
			if View.IsNumeric(number) {
				n = StringToInteger(number)
				break
			} else {
				fmt.Println("Bitte geben Sie eine gültige Eingabe [1-4] ein.")
				continue
			}
		}

		if n > 4 && n < 0 {
			continue
		} else {
			switch n {
			case 1:
				status = "Backlog"
			case 2:
				status = "To-Do"
			case 3:
				status = "Ongoing"
			case 4:
				status = "Done"
			}
		}

		newTask := Model.Task{
			Task:      input,
			Status:    status,
			Timestamp: time.Now().Format(time.RFC822Z),
		}

		tasks = append(tasks, newTask)
	}

	if input == "q" && cap(tasks) == 0 {
		View.Quit()
	} else {
		Model.SaveTask(&tasks)
		//View.Quit()
		LoadTaskAll()
	}
}

// EditTask → Die Aufgabe bearbeiten
func EditTask() {
	View.Clear()

	var (
		input string
		tasks []Model.Task
		id    int
	)

	tasks = *Model.TasksTable()

	fmt.Println("############################################################################")

	Model.AllList()

	fmt.Println("############################################################################")

	fmt.Println("Welche Aufgabe möchten Sie bearbeiten?")
	fmt.Println("Bitte geben Sie einen Nummer ein oder 'q' ein, um zu beenden:")

	for {
		var (
			newTask   string
			newStatus string
		)

		input = View.AskForInput()

		if input == "q" {
			break
		} else {
			id = StringToInteger(input)

			if id > 0 && id <= len(tasks) {
				fmt.Printf("Bitte geben Sie die neue Beschreibung für Aufgabe %d ein:\n", id)
				newTask = View.AskForInput()

				fmt.Print("Klassifizierung:\n[1] Backlog\n[2] To-Do\n[3] Ongoing\n[4] Done\nBitte wählen Sie die Status:\n")

				newStatus = View.AskForInput()

				if newStatus == "" {
					newStatus = ""
				} else {
					var n int

					if View.IsNumeric(newStatus) {
						n = StringToInteger(newStatus)
					} else {
						fmt.Println("Bitte geben Sie eine gültige Eingabe ein.")
						continue
					}

					if n > 4 && n < 0 {
						fmt.Println("Bitte geben Sie eine gültige Eingabe ein.")
						continue
					} else {
						switch n {
						case 1:
							newStatus = "Backlog"
						case 2:
							newStatus = "To-Do"
						case 3:
							newStatus = "Ongoing"
						case 4:
							newStatus = "Done"
						}
					}
				}

				newId := id - 1

				if newStatus == "" && newTask != "" {
					tasks[newId] = Model.Task{
						Task:      newTask,
						Status:    tasks[newId].Status,
						Timestamp: time.Now().Format(time.RFC822Z),
					}
				} else if newTask == "" && newStatus != "" {
					tasks[newId] = Model.Task{
						Task:      tasks[newId].Task,
						Status:    newStatus,
						Timestamp: time.Now().Format(time.RFC822Z),
					}
				} else if newTask != "" && newStatus != "" {
					tasks[newId] = Model.Task{
						Task:      newTask,
						Status:    newStatus,
						Timestamp: time.Now().Format(time.RFC822Z),
					}
				} else {
					fmt.Println("Aufgabe wurde unverändert.")
				}
			} else {
				fmt.Println("Bitte geben Sie eine gültige Eingabe ein.")
				continue
			}
		}
		break
	}

	if input != "q" {
		Model.UpdateTask(&tasks)
		fmt.Printf(
			"Aufgabe %d wurde aktualisiert.\n",
			id,
		)
		//View.Quit()
		LoadTaskAll()
	} else {
		View.Quit()
	}
}

// RemoveTask → Die Aufgabe löschen
func RemoveTask() {
	View.Clear()

	var (
		id    string
		tasks []Model.Task
	)

	fmt.Println("############################################################################")

	Model.AllList()

	fmt.Println("############################################################################")

	tasks = *Model.TasksTable()

	for {
		var (
			remove int
		)

		fmt.Println("Welche Aufgabe möchten Sie entfernen?")
		fmt.Print("Bitte geben Sie einen Nummer ein oder 'q' ein, um zu beenden:")

		id = View.AskForInput()

		if id == "q" {
			break
		} else {
			remove = StringToInteger(id)

			if remove > 0 && remove <= len(tasks) {
				tasks = append(tasks[:remove-1], tasks[remove:]...)
				fmt.Printf("Aufgabe %d wurde entfernt\n", remove)
				break
			} else {
				fmt.Println("Bitte geben Sie eine gültige Eingabe ein.")
				continue
			}
		}
	}

	if id == "q" {
		View.Quit()
	} else {
		Model.UpdateTask(&tasks)
		//View.Quit()
		LoadTaskAll()
	}
}

// LoadTaskAll → Die gesamte Listen auf dem Fenster anzeigen und die Option bearbeiten
func LoadTaskAll() {
	View.Clear()

	var check int

	fmt.Println("Alle Aufgaben anzeigen:")
	fmt.Println("############################################################################")

	Model.AllList()

	fmt.Println("############################################################################")
	fmt.Println("\n[1] Neue Aufgabe erstellen, [2] Bearbeiten, [3] Löschen, [4] Beenden")

	for {
		fmt.Print("Wählen Sie eine Option vor Tabelle: ")
		input := View.AskForInput()

		if View.IsNumeric(input) {
			check = StringToInteger(input)
		} else {
			fmt.Println("Bitte geben Sie eine gültige Eingabe ein.")
			continue
		}

		switch check {
		case 1, 2, 3:
			SwitchChoose(check)
			break
		case 4:
			View.Menu()
			break
		default:
			fmt.Println("Bitte geben Sie eine gültige Eingabe ein.")
			continue
		}
		break
	}
}

// StringToInteger → Die Zeichenketten werden in die ganze Zahl umgewandelt
func StringToInteger(input string) int {
	newInput, _ := strconv.Atoi(input)
	return newInput
}
