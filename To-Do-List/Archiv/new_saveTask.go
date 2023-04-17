package Archiv

/*func saveTask(tasks []Model.Task) {

	newTasks := tasks

	file, err := os.OpenFile("tasks.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Fehler beim Erstellen der Datei:", err)
		Menu()
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Comma = '\t'

	fileStat, err := file.Stat()
	if err != nil {
		fmt.Println("Fehler beim Abrufen der Dateigrösse:", err)
		return
	}

	fileSize := fileStat.Size()

	if fileSize == 0 {
		header := []string{"", "", "", ""}
		header[0] = "ID"
		header[1] = "Tasks"
		header[2] = "Status"
		header[3] = "Timestamp"
		if err := writer.Write(header); err != nil {
			fmt.Println("Fehler beim Schreiben der Kopfzeile:", err)
			Menu()
		}
	}

	nextID, err := Model.getNextTaskID()
	if err != nil {
		fmt.Println("Fehler beim Ermitteln der nächsten ID:", err)
		Menu()
	} else if nextID == 0 {
		nextID = 1
	}

	for _, task := range newTasks {
		taskData := []string{"", "", "", ""}
		taskData[0] = strconv.Itoa(nextID)
		taskData[1] = task.Task
		taskData[2] = task.Status
		taskData[3] = task.Timestamp

		if err := writer.Write(taskData); err != nil {
			fmt.Println("Fehler beim Schreiben der Datei:", err)
			Menu()
		}
		nextID++

		writer.Flush()
	}

	fmt.Println("Die neue Aufgabe wurde hinzugefügt und in der Datei gespeichert.")
}*/
