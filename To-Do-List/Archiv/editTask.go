package Archiv

/*func editTask() {
	var (
		id    string
		tasks = TasksTable()
	)

	fmt.Print("\033[H\033[2J")
	fmt.Println("Welche Aufgabe möchten Sie bearbeiten?\n############################################################################")

	AllList()

	fmt.Println("############################################################################")

	fmt.Println("Bitte geben Sie einen Nummer ein oder 'q' ein, um zu beenden:")

	for {
		if _, err := fmt.Scanln(&id); err != nil {
			fmt.Println("Ungültige Auswahl.")
			continue
		} else if id == "q" {
			fmt.Println("Dies ist hier beendet.")
			fmt.Printf("Lädt...")
			time.Sleep(3 * time.Second)
			Menu()
		} else {
			id, err := strconv.Atoi(id)

			if err != nil {
				fmt.Println("Ungültige Auswahl.")
				continue
			}

			if id > 0 && id <= len(tasks) {
				fmt.Printf("Bitte geben Sie die neue Beschreibung für Aufgabe %d ein:\n", id)
				newTask, _ := bufio.NewReader(os.Stdin).ReadString('\n')
				newTask = strings.TrimSpace(newTask)

				fmt.Print("Fertigungssteuerung:\n[1] Backlog\n[2] To-Do\n[3] Ongoing\n[4] Done\nBitte wählen Sie die Status:\n")

				newStatus, _ := bufio.NewReader(os.Stdin).ReadString('\n')
				newStatus = strings.TrimSpace(newStatus)

				if err != nil {
					fmt.Println("Fehler beim Lesen der Eingabe:", err)
					continue
				}

				if newStatus == "" {
					newStatus = ""
				} else {
					n, err := strconv.Atoi(newStatus)

					if err != nil && n > 4 && n < 0 {
						fmt.Println("Fehler beim Lesen der Eingabe:", err)
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

				newStatus = strings.TrimSpace(newStatus)

				newId := id - 1

				if newStatus == "" && newTask != "" {
					tasks[newId] = Model.Task{
						Task:      newTask,
						Status:    tasks[newId].Status,
						Timestamp: time.Now().Format(time.RFC822Z),
					}
					break
				} else if newTask == "" && newStatus != "" {
					tasks[newId] = Model.Task{
						Task:      tasks[newId].Task,
						Status:    newStatus,
						Timestamp: time.Now().Format(time.RFC822Z),
					}
					break
				} else if newTask != "" && newStatus != "" {
					tasks[newId] = Model.Task{
						Task:      newTask,
						Status:    newStatus,
						Timestamp: time.Now().Format(time.RFC822Z),
					}
					break
				} else {
					fmt.Println("Aufgabe wurde unverändert.\n")
					break
				}
			} else {
				fmt.Println("Ungültige Auswahl.")
				continue
			}
		}
	}

	newSaveTask(tasks)
	fmt.Printf(
		"Aufgabe %s wurde aktualisiert.\n",
		id,
	)
	fmt.Printf("Lädt...")
	time.Sleep(3 * time.Second)
	Menu()
}*/
