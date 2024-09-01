package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/eterline/echidna/internal/gotify"
	"github.com/eterline/echidna/internal/settings"
	"github.com/gorilla/mux"
)

type Server struct {
	Router   *mux.Router
	Settings settings.Config
}

func New(c settings.Config) *Server {
	var s Server
	s.Router = mux.NewRouter()
	s.Settings = c
	s.Router.Use(s.catcher)
	s.Router.HandleFunc("/", base)
	return &s
}

func base(w http.ResponseWriter, r *http.Request) {
	text := fmt.Sprintf("Hi %s, you have been burned", r.RemoteAddr)
	fmt.Fprintln(w, text)
}

func (s *Server) catcher(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s | %s < catched.", r.RemoteAddr, r.RequestURI)
		go gotify.SendMessage(r, s.Settings)
		next.ServeHTTP(w, r)
	})
}
