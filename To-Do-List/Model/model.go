package Model

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type Task struct {
	ID        int
	Task      string
	Status    string
	Timestamp string
}

type Account struct {
	Username string
	Password string
}

// CheckUser → Hier wird ein Benutzername und Passwort überprüfen
func CheckUser(user string, password string) bool {
	file, err := os.Open("users.csv")

	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	reader := NewReaderForCSV(file)
	records, err := reader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records {
		if record[0] == user && record[1] == password {
			return true
		}
	}

	return false
}

// RegisterUser → Hier wird ein neuer Benutzername und Passwort erfasst
func RegisterUser(signUp *[]Account) {
	var (
		file *os.File
		err  error
	)

	if file, err = os.OpenFile("users.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644); Error(err) == true {
		log.Fatal("Fehler beim Erstellen der Datei:", err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	writer := NewWriterForCSV(file)
	defer writer.Flush()

	for _, record := range *signUp {
		taskData := []string{"", ""}
		taskData[0] = record.Username
		taskData[1] = record.Password

		err = writer.Write(taskData)

		if err != nil {
			log.Fatalln(err)
		}
	}

	fmt.Println("Registrierung erfolgreich abgeschlossen")
	return
}

// TasksTable → Die Tabelle wird als .csv-Datei geladen und für die Aktualisierung verwendet
func TasksTable() *[]Task {
	var tasks []Task

	file, err := os.Open("tasks.csv")

	if Error(err) == true {
		log.Fatalln(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	reader := NewReaderForCSV(file)

	for {
		row, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln(err)
		}

		id, err := strconv.Atoi(row[0])

		//Fehler beim Einfügen des Titels von ID ('no string')
		if err != nil {
			continue
		}

		task := Task{
			ID:        id,
			Task:      row[1],
			Status:    row[2],
			Timestamp: row[3],
		}

		tasks = append(tasks, task)
	}

	return &tasks
}

// SaveTask → Hier wird eine weitere Zeile hinzugefügt oder eine neue Tabelle erfasst
func SaveTask(tasks *[]Task) {
	var (
		file   *os.File
		err    error
		nextID int
	)

	if file, err = os.OpenFile("tasks.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644); Error(err) == true {
		log.Fatal("Fehler beim Erstellen der Datei:", err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	writer := NewWriterForCSV(file)

	fileStat, err := file.Stat()

	if Error(err) == true {
		log.Fatalln("Fehler beim Abrufen der Dateigrösse:", err)
	}

	fileSize := fileStat.Size()

	if fileSize == 0 {
		header := []string{"", "", "", ""}
		header[0] = "ID"
		header[1] = "Tasks"
		header[2] = "Status"
		header[3] = "Timestamp"
		if err := writer.Write(header); err != nil {
			log.Fatalln("Fehler beim Schreiben der Kopfzeile:", err)
		}
	}

	if nextID, err = getNextTaskID(); Error(err) == true {
		log.Fatalln("Fehler beim Ermitteln der nächsten ID:", err)
	} else if nextID == 0 {
		nextID = 1
	}

	for _, task := range *tasks {
		taskData := []string{"", "", "", ""}
		taskData[0] = IntegerToString(nextID)
		taskData[1] = task.Task
		taskData[2] = task.Status
		taskData[3] = task.Timestamp

		if err := writer.Write(taskData); Error(err) == true {
			log.Fatalln("Fehler beim Schreiben der Datei:", err)
		}
		nextID++

		// Flushen der Buffer und Speichern der Daten
		writer.Flush()

		if err := writer.Error(); Error(err) == true {
			log.Fatalln("Fehler beim Schreiben der Datei:", err)
		}
	}

	fmt.Println("Die neue Aufgabe wurde hinzugefügt und in der Datei gespeichert.")
}

// UpdateTask → Hier können Zeilen geändert und gelöscht werden
func UpdateTask(tasks *[]Task) {
	var (
		file *os.File
		err  error
	)

	if file, err = os.OpenFile("tasks.csv", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644); Error(err) == true {
		log.Fatalln("Fehler beim Erstellen der Datei:", err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	writer := NewWriterForCSV(file)

	fileStat, err := file.Stat()

	if Error(err) == true {
		log.Fatalln("Fehler beim Abrufen der Dateigrösse:", err)
	}

	fileSize := fileStat.Size()

	if fileSize == 0 {
		header := []string{"", "", "", ""}
		header[0] = "ID"
		header[1] = "Tasks"
		header[2] = "Status"
		header[3] = "Timestamp"
		if err := writer.Write(header); Error(err) == true {
			log.Fatalln("Fehler beim Schreiben der Kopfzeile:", err)
		}
	}

	nextID, err := getNextTaskID()
	if Error(err) == true {
		log.Fatalln("Fehler beim Ermitteln der nächsten ID:", err)
	} else if nextID == 0 {
		nextID = 1
	}

	for _, task := range *tasks {
		taskData := []string{"", "", "", ""}
		taskData[0] = IntegerToString(nextID)
		taskData[1] = task.Task
		taskData[2] = task.Status
		taskData[3] = task.Timestamp

		if err := writer.Write(taskData); Error(err) == true {
			log.Fatalln("Fehler beim Schreiben der Datei:", err)
		}
		nextID++

		writer.Flush()

		if err := writer.Error(); Error(err) == true {
			log.Fatalln("Fehler beim Schreiben der Datei:", err)
		}
	}

	fmt.Println("Die neue Aufgabe wurde hinzugefügt und in der Datei gespeichert.")
	return
}

// AllList → Die gesamte Tabelle wird auf View angezeigt
func AllList() {
	var (
		file *os.File
		err  error
	)

	if file, err = os.Open("tasks.csv"); Error(err) == true {
		log.Fatalln("Fehler beim Öffnen der Datei:", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	reader := NewReaderForCSV(file)

	row, err := reader.ReadAll()

	if Error(err) == true {
		log.Fatalln("Fehler beim Lesen der Datei:", err)
	}

	for _, task := range row {
		id := task[0]
		tasks := task[1]
		status := task[2]

		fmt.Printf("%s\t%-25s\t%s\n", id, tasks, status)
	}
}

// NewReaderForCSV → Datei von .csv wird auf der Tabelle erzeugen
func NewReaderForCSV(file *os.File) *csv.Reader {
	reader := csv.NewReader(file)
	reader.Comma = '\t'

	return reader
}

// NewWriterForCSV → .csv wird aus der Standard-Datenströme auf der Datei abgeleitet
func NewWriterForCSV(file *os.File) *csv.Writer {
	writer := csv.NewWriter(file)
	writer.Comma = '\t'

	return writer
}

// Error → Überprüfen der Fehlerausgabe
func Error(err error) bool {
	if err != nil {
		return true
	}
	return false
}

// IntegerToString → Die ganze Zahl werden in die Zeichenketten umgewandelt
func IntegerToString(input int) string {
	return strconv.Itoa(input)
}

// getNextTaskID → Um die Identifikatoren auf jede Zeile automatisch zu erfassen
func getNextTaskID() (int, error) {
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
}
