package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type User struct {
	Username string
	Password string
}

var users = map[string]User{
	"john": {
		Username: "john",
		Password: "password123",
	},
	"jane": {
		Username: "jane",
		Password: "qwerty",
	},
}

func main() {
	http.HandleFunc("/", loginHandler)
	http.HandleFunc("/login", authenticateHandler)
	log.Fatal(http.ListenAndServe(":8090", nil))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("login.html"))

	if err := tmpl.Execute(w, nil); err != nil {
		log.Println(err)
	}
}

func authenticateHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	user, ok := users[username]
	if !ok || user.Password != password {
		fmt.Fprintln(w, "Invalid username or password")
		return
	}

	// Successful login
	fmt.Fprintf(w, "Welcome, %s!", username)
}
