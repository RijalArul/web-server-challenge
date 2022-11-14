package main

import (
	"fmt"
	"net/http"
	"text/template"
	"web-server-2-go/biodata"
)

func setCookie(w http.ResponseWriter, r *http.Request, name string) {
	cookie := new(http.Cookie)
	cookie.Name = "Name"
	cookie.Value = name
	cookie.Path = "/"
	http.SetCookie(w, cookie)
}

func main() {
	bio := biodata.Biodata{}
	mux := http.NewServeMux()
	t := template.Must(template.ParseGlob("./templates/*.gohtml"))
	listNama := bio.ListNama()
	LoginPage := func(w http.ResponseWriter, r *http.Request) {
		_, errC := r.Cookie("Name")
		err := r.ParseForm()

		username := r.PostFormValue("username")
		password := r.PostFormValue("password")

		if err != nil {
			panic(err)
		}

		if username == bio.FilterNama(username) && password == "Abc123" || errC == nil {
			setCookie(w, r, bio.FilterNama(username))
			http.Redirect(w, r, "http://localhost:8080/biodata", http.StatusSeeOther)
		}

		t.ExecuteTemplate(w, "login-page.gohtml", nil)

	}

	BioPage := func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("Name")

		if err != nil {
			http.Redirect(w, r, "http://localhost:8080/login", http.StatusSeeOther)
		} else {
			if len(listNama) > 0 && cookie.Value != "" {
				if bio.FilterNama(cookie.Value) != "" {
					bio.FilterNama(cookie.Value)
					bio.GenerateID()
					bio.GenerateAddress()
					bio.GenerateJob()
					bio.GenerateReason()
				}
			}
		}
		t.ExecuteTemplate(w, "bio-page.gohtml", bio)

	}

	LogOut := func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("Name")
		cookie.MaxAge = -1
		fmt.Println(cookie, err)
		http.SetCookie(w, cookie)
		http.Redirect(w, r, "/login", http.StatusFound)

	}

	mux.HandleFunc("/login", LoginPage)
	mux.HandleFunc("/biodata", BioPage)
	mux.HandleFunc("/logout", LogOut)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}

	fmt.Println("starting web server at http://localhost:8080/")

}
