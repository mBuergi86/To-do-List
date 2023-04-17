package Model

/*func getNextTaskID() (int, error) {
	file, err := os.Open("tasks.csv")
	if err != nil {
		return 0, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	rows, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println("Fehler beim Lesen der Datei:", err)
		return 0, nil
	}

	return len(rows), nil
}*/
