package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"sync"
	"text/template"

	"github.com/AlekseiKanash/golang-course/lesson_10/store/src/pgs"
	"github.com/gorilla/mux"
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
	router    *mux.Router
	wg        *sync.WaitGroup
	isInit    bool
}

func (s *Server) Init() {
	s.router = mux.NewRouter()
	s.router.HandleFunc("/api/v0/fetch/{name}", s.userHandler).Methods(http.MethodGet)
	s.router.HandleFunc("/api/v0/city/{name}", s.cityHandler).Methods(http.MethodGet)
	s.router.HandleFunc("/api/v0/list", s.listHandler).Methods(http.MethodGet)
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

	s.server = &http.Server{Addr: s.Addr, Handler: s.router}

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

func retJson(w http.ResponseWriter, data interface{}) {
	str, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("%v", err)
	}
	w.Header().Add("Content-Type", "application/json")
	if _, err := w.Write(str); err != nil {
		fmt.Printf("%v", err)
	}
}

func (s *Server) cityHandler(w http.ResponseWriter, r *http.Request) {
	city := mux.Vars(r)["name"]
	if city == "" {
		fmt.Println("Wrong Request: city var does not exist")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	weather := pgs.GetWeather(city)

	retJson(w, weather)
}

func (s *Server) listHandler(w http.ResponseWriter, r *http.Request) {
	weathers := pgs.ListWeather()
	retJson(w, weathers)
}

func (s *Server) userHandler(w http.ResponseWriter, r *http.Request) {
	city := mux.Vars(r)["name"]
	if city == "" {
		fmt.Println("Wrong Request: city var does not exist")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	switch r.RequestURI {
	case "/":
		{
			if r.Method == http.MethodPost {
				city := r.PostFormValue("City")
				pgs.SaveWeather(city)
			}
		}
	}

	cookie, err := r.Cookie("token")
	if err != nil {
		cookie = &http.Cookie{
			Name:     "token",
			Value:    "",
			HttpOnly: false,
		}
		http.SetCookie(w, cookie)
	}
	renderTemplate(w, "user", cookie)
}
