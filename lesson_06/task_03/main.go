package main

import (
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

var validPath = regexp.MustCompile("^/(edit|save|view|user|saveUser)/([a-zA-Z0-9]+)$")
var templates = template.Must(template.ParseFiles("user.html"))

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("invalid Page Title")
	}
	return m[2], nil // The title is the second subexpression.
}

func loadUserPage(title string) ([]byte, error) {
	filename := title + ".html"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p interface{}) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		cookie = &http.Cookie{
			Name:     "token",
			Value:    "",
			HttpOnly: false,
		}
		http.SetCookie(w, cookie)
	}

	switch r.RequestURI {
	case "/":
		{
			if r.Method == http.MethodPost {
				name := r.PostFormValue("name")
				address := r.PostFormValue("address")
				cookie.Value = fmt.Sprintf("%s:%s", name, address)
				http.SetCookie(w, cookie)
			}
		}
	default:
		{
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}

	renderTemplate(w, "user", cookie)
}

func main() {
	http.HandleFunc("/", userHandler)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
