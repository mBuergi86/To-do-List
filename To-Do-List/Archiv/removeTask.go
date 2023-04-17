package Archiv

/*func removeTask() {
	var (
		id       string
		tasks    = TasksTable()
		newTasks = Model.Tasks
	)

	fmt.Print("\033[H\033[2J")
	fmt.Println("Welche Aufgabe möchten Sie entfernen?\n############################################################################")

	AllList()

	fmt.Println("############################################################################")

	for {
		fmt.Println("Bitte geben Sie einen Nummer ein oder 'q' ein, um zu beenden:")

		if _, err := fmt.Scanln(&id); err != nil {
			fmt.Println("Ungültige Auswahl.")
			continue
		} else if id == "q" {
			fmt.Println("Dies ist hier beendet.")
			fmt.Printf("Lädt...")
			time.Sleep(3 * time.Second)
			Menu()
		} else {
			remove, err := strconv.Atoi(id)

			if err != nil {
				fmt.Println("Ungültige Auswahl.")
				continue
			}

			if remove > 0 && remove <= len(tasks) {
				newTasks = append(tasks[:remove-1], tasks[remove:]...)
				fmt.Printf("Aufgabe %d wurde entfernt\n", remove)
				break
			} else {
				fmt.Println("Ungültige Auswahl.")
				continue
			}
		}
	}

	newSaveTask(newTasks)

	fmt.Printf("Lädt...")
	time.Sleep(3 * time.Second)
	Menu()
}*/
