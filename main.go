package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Biodata struct {
	Nama, Email, Alamat, Pekerjaan, Alasan string
}

var PORT = ":3333"

func main() {
	http.HandleFunc("/", renderTemplate)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)

	fmt.Println("Application is listening on port", PORT)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatal("Error:", err)
		return
	}
}

func renderTemplate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var template = template.Must(template.New("form").ParseFiles("login.html"))
		var err = template.Execute(w, nil)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}
}

func login(w http.ResponseWriter, r *http.Request) {

}

func logout(w http.ResponseWriter, r *http.Request) {

}
