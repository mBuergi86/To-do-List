package Archiv

/*func TasksTable() []Model.Task {
	file, err := os.Open("tasks.csv")

	if err != nil {
		panic(err)
	}

	reader := csv.NewReader(file)
	reader.Comma = '\t'

	if err != nil {
		fmt.Println("Fehler beim Abrufen der Dateigrösse:", err)
		Menu()
	}

	tasks := Model.Tasks

	for {
		row, err := reader.Read()

		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		id, err := strconv.Atoi(row[0])

		if err != nil {
			continue
		}

		task := Model.Task{
			ID:        id,
			Task:      row[1],
			Status:    row[2],
			Timestamp: row[3],
		}

		tasks = append(tasks, task)
	}

	return tasks
}*/
