package View

import (
	"bufio"
	"fmt"
	"golang.org/x/term"
	"log"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"syscall"
)

// Menu → Startseite
func Menu() {
	Clear()

	fmt.Println("########Willkommen bei der To-do-Liste#########")
	fmt.Println()
	fmt.Println("[1] Neue Aufgabe hinzufügen")
	fmt.Println("[2] Bestehende Aufgabe bearbeiten")
	fmt.Println("[3] Bestehende Aufgabe entfernen")
	fmt.Println("[4] Alle Aufgaben anzeigen")
	fmt.Println("[5] Beenden")
	fmt.Println()
	fmt.Print("Wählen Sie eine Option: ")
}

// AccountMenu → Sicherheitsabfrage
func AccountMenu() {
	Clear()

	fmt.Println("########Zugangssicherheit - To-do-Liste########")
	fmt.Println()
	fmt.Println("[1] Login")
	fmt.Println("[2] Neue Registrierung")
	fmt.Println("[3] Beenden")
	fmt.Println()
	fmt.Print("Wählen Sie eine Option: ")
}

// AskForInput → Standard-Datenströme für die Eingabe
func AskForInput() string {
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')

	if err != nil {
		log.Fatalln(err)
	}

	input = strings.TrimSpace(input)

	return input
}

// PasswordOfInput → Das Passwort wird ausblenden
func PasswordOfInput() string {
	byteInput, err := term.ReadPassword(int(syscall.Stdin))

	if err != nil {
		log.Fatal(err)
	}

	input := strings.TrimSpace(string(byteInput))

	return input
}

// Clear → Die Eingabeaufforderung wurde bereinigt
func Clear() {
	switch runtime.GOOS {
	case "darwin":
		runCmd("clear")
	case "linux":
		runCmd("clear")
	case "windows":
		runCmd("cmd", "/c", "cls")
	default:
		runCmd("clear")
		//fmt.Print("\033[H\033[2J")
	}
}

// Die Befehle von Clear ausführen
func runCmd(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		return
	}
}

// Quit → Beenden und auf die Startseite zurückkehren
func Quit() {
	//fmt.Println("Dies ist hier beendet.")
	//time.Sleep(3 * time.Second)
	Menu()
}

// IsNumeric -> Die Eingabe prüft, ob die arabischen Ziffern als String vorliegen
func IsNumeric(input string) bool {
	return regexp.MustCompile(`^[0-9]+$`).MatchString(input)
}
