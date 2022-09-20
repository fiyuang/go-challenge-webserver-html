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
	if r.Method == "POST" {
		var template = template.Must(template.New("home").ParseFiles("index.html"))
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var email = r.FormValue("email")
		var keyEmail int

		emails := []string{"fitri@mail.com", "ayu@mail.com", "anggraini@mail.com", "fiyuang@mail.com", "fifiyuu@mail.com"}
		result := getDataPeserta(emails)
		// fmt.Println(result)

		// check email exist
		for key, value := range emails {
			if email == value {
				keyEmail = key
			}
		}

		// deliver data to template
		for key, value := range result {
			if keyEmail == key {
				var data = map[string]string{
					"email":     email,
					"nama":      value.Nama,
					"alamat":    value.Alamat,
					"pekerjaan": value.Pekerjaan,
					"alasan":    value.Alasan,
				}
				var err = template.Execute(w, data)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}
		}

		return
	}
}

func logout(w http.ResponseWriter, r *http.Request) {

}

func getDataPeserta(emails []string) []Biodata {
	var dataPeserta = []Biodata{
		{
			Nama:      "Fitri",
			Alamat:    "Jl. Lorem",
			Pekerjaan: "Backend",
			Alasan:    "Alasan Fitri",
		},
		{
			Nama:      "Ayu",
			Alamat:    "Jl. Ipsum",
			Pekerjaan: "Frontend",
			Alasan:    "Alasan Ayu",
		},
		{
			Nama:      "Anggraini",
			Alamat:    "Jl. Dolor",
			Pekerjaan: "Fullstack",
			Alasan:    "Alasan Anggraini",
		},
		{
			Nama:      "Fiyuang",
			Alamat:    "Jl. Sit",
			Pekerjaan: "DevOps",
			Alasan:    "Alasan Fiyuang",
		},
		{
			Nama:      "Fifiyuu",
			Alamat:    "Jl. Amet",
			Pekerjaan: "Mobile Dev",
			Alasan:    "Alasan Fifiyuu",
		},
	}

	response := make([]Biodata, 0)
	var data Biodata

	for key, value := range emails {
		data.Email = value
		data.Nama = dataPeserta[key].Nama
		data.Alamat = dataPeserta[key].Alamat
		data.Pekerjaan = dataPeserta[key].Pekerjaan
		data.Alasan = dataPeserta[key].Alasan
		response = append(response, data)
	}

	return response
}
