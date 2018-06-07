package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Server responds to http requests.
type Server struct {
	router *mux.Router
}

type route struct {
	method  string
	path    string
	handler http.HandlerFunc
}

// New creates a server that responds to http requests.
func New() *Server {
	s := &Server{router: mux.NewRouter()}

	rr := []route{
		{method: "GET", path: "/", handler: Home},
		{method: "GET", path: "/settings", handler: Settings},
		{method: "GET", path: "/user/{username}", handler: Profile},
	}

	s.router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	for _, r := range rr {
		s.router.HandleFunc(r.path, r.handler).Methods(r.method)
	}

	return s
}

// ServeHTTP responds to http requests by delegating to the server.router.
func (server Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	server.router.ServeHTTP(w, r)
}

// Home responds with the root index.html.
func Home(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir("./static/html/")).ServeHTTP(w, r)
}