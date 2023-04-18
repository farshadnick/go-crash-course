package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

type Person struct {
	Name string
	Date time.Time
}

var people []Person
var filename string = "people.txt"

func main() {
	http.HandleFunc("/add-person", addPersonHandler)
	http.HandleFunc("/check-today", checkTodayHandler)
	http.HandleFunc("/check", checkHandler)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

func addPersonHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	dateStr := r.FormValue("date")

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		http.Error(w, "Invalid date format. Use yyyy-mm-dd", http.StatusBadRequest)
		return
	}

	person := Person{Name: name, Date: date}
	people = append(people, person)

	// save the person to the file
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintf(w, "Failed to open file: %v", err)
		return
	}
	defer f.Close()

	line := fmt.Sprintf("%s %s\n", name, date.Format("2006-01-02"))
	if _, err := f.WriteString(line); err != nil {
		fmt.Fprintf(w, "Failed to write to file: %v", err)
		return
	}

	fmt.Fprintf(w, "Added person: %s, %s\n", name, dateStr)
}

func checkTodayHandler(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	todayMonth := now.Month()
	todayDay := now.Day()

	for _, person := range people {
		personMonth := person.Date.Month()
		personDay := person.Date.Day()

		if personMonth == todayMonth && personDay == todayDay {
			fmt.Fprintf(w, "It's today! Happy birthday %s!\n", person.Name)
		}
	}
}

func checkHandler(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(w, "Failed to open file: %v", err)
		return
	}
	defer f.Close()

	// read the contents of the file
	var contents string
	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			fmt.Fprintf(w, "Failed to read file: %v", err)
			return
		}
		contents += string(buf[:n])
	}

	fmt.Fprintf(w, "Contents of %s:\n%s", filename, contents)
}
