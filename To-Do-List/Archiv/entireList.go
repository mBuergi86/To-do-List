package Archiv

/*func AllList() {
	file, err := os.Open("tasks.csv")

	if err != nil {
		fmt.Println("Fehler beim Öffnen der Datei:", err)
		Menu()
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = '\t'

	row, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Fehler beim Lesen der Datei:", err)
		Menu()
	}

	for _, task := range row {
		id := task[0]
		tasks := task[1]
		status := task[2]

		fmt.Printf("%s\t%-25s\t%s\n", id, tasks, status)
	}
}*/
