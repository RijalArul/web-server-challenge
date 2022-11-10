package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	dataEmail := []string{"Rijalarul7599@gmail.com"}
	data := map[string]string{}

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "POST" {
			usernameForm := r.FormValue("username")
			password := r.FormValue("password")
			for i := 0; i < len(dataEmail); i++ {
				if usernameForm == dataEmail[i] && len(password) > 0 {
					username := dataEmail[i]
					data = map[string]string{
						"username": username,
						"alasan":   "Alasan " + username,
					}

					http.Redirect(w, r, "http://localhost:8080/", http.StatusSeeOther)
				}
			}
		}

		var t, err = template.ParseFiles("login.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(w, data)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if len(data) == 0 {
			http.Redirect(w, r, "http://localhost:8080/login", http.StatusSeeOther)
		} else {
			var t, err = template.ParseFiles("template.html")
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			t.Execute(w, data)
		}

	})

	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "POST" {
			data = map[string]string{}
			http.Redirect(w, r, "http://localhost:8080/login", http.StatusSeeOther)

			var t, err = template.ParseFiles("template.html")
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			fmt.Println(data)

			t.Execute(w, data)
		}

	})

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)

}
