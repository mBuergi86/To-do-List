package Archiv

/*func addTask() {
	tasks := Model.Tasks
	fmt.Print("\033[H\033[2J")

	for {
		reader := bufio.NewReader(os.Stdin)

		fmt.Println("Bitte geben Sie die neue Aufgabe ein oder geben Sie 'q' ein, um zu beenden:")
		task, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Fehler beim Lesen der Eingabe:", err)
			continue
		}

		task = strings.TrimSpace(task)

		if task == "q" {
			fmt.Println("Dies ist hier beendet.")
			break
		}

		fmt.Print("Fertigungssteuerung:\n[1] Backlog\n[2] To-Do\n[3] Ongoing\n[4] Done\nBitte wählen Sie die Status:\n")

		number, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Fehler beim Lesen der Eingabe:", err)
			continue
		}

		number = strings.TrimSpace(number)
		n, err := strconv.Atoi(number)
		var status string

		if err != nil && n > 4 && n < 0 {
			fmt.Println("Fehler beim Lesen der Eingabe:", err)
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
			Task:      task,
			Status:    status,
			Timestamp: time.Now().Format(time.RFC822Z),
		}

		tasks = append(tasks, newTask)
	}

	saveTask(tasks)
	fmt.Printf("Lädt...")

	time.Sleep(5 * time.Second)
	Model.Count = 5
	  for Model.Count > 0 {
	  	time.Sleep(1 * time.Second)
	  	Model.Count--
	  }
	Menu()
}*/
