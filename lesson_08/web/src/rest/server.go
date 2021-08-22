package rest

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"sync"
	"text/template"
)

var validPath = regexp.MustCompile("^/(user)/([a-zA-Z0-9]+)$")
var templates = template.Must(template.ParseFiles("user.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p interface{}) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type Server struct {
	IsRunning bool
	Addr      string
	server    *http.Server
	wg        *sync.WaitGroup
	isInit    bool
}

func (s *Server) Init() {
	http.HandleFunc("/", s.userHandler)
}

func (s *Server) Start() {
	if s.IsRunning {
		fmt.Println("Already running.")
		return
	}

	if nil == s.wg {
		s.wg = &sync.WaitGroup{}
	}

	s.wg.Add(1)

	s.server = &http.Server{Addr: s.Addr}

	go func(wg *sync.WaitGroup) {
		defer fmt.Println("done")
		defer wg.Done() // let main know we are done cleaning up

		// always returns error. ErrServerClosed on graceful close
		s.IsRunning = true
		fmt.Printf("Running HTTP Listen server on %s\n", s.Addr)
		err := s.server.ListenAndServe()
		if err != http.ErrServerClosed {
			// unexpected error. port in use?
			log.Fatalf("ListenAndServe(): %v", err)
		}
		fmt.Printf("Server is stopped. %v\n", err)
		s.IsRunning = false

	}(s.wg)
}

func (s *Server) Stop() {
	if s.IsRunning {
		if err := s.server.Shutdown(context.TODO()); err != nil {
			panic(err) // failure/timeout shutting down the server gracefully
		}

		// wait for goroutine started in startHttpServer() to stop
		s.wg.Wait()
	}
}

func (s *Server) userHandler(w http.ResponseWriter, r *http.Request) {
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
