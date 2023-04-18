package main

import (
	"fmt"
	"net/http"
)


func main() {
	http.HandleFunc("/w", handler)
  http.ListenAndServe(":8080", nil)
	fmt.Println("Server is running on port 8080")
  
func handler(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()

    name := r.Form.Get("name")
    email := r.Form.Get("email")

    fmt.Fprintf(w, "Name: %s\n", name)
    fmt.Fprintf(w, "Email: %s\n", email)
  }
}
