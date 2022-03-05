package server

import (
	"net/http"

	"github.com/EgorSkurihin/url-shortener/store"
	"github.com/gorilla/mux"
)

type Server struct {
	store  *store.Store
	router mux.Router
}

func New(dsn string) *Server {
	s := &Server{
		store:  &store.Store{DSN: dsn},
		router: *mux.NewRouter(),
	}
	s.configureRouter()
	return s
}

func (srv *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	srv.router.ServeHTTP(w, r)
}

func (srv *Server) configureRouter() {
	srv.router.HandleFunc("/", srv.index).Methods("GET")
	srv.router.HandleFunc("/create", srv.createURL).Methods("POST")
	srv.router.HandleFunc("/{URL}", srv.showURL).Methods("GET")
}

func (srv *Server) Run(addr string) error {
	if err := srv.store.Open(); err != nil {
		return err
	}
	defer srv.store.Close()
	return http.ListenAndServe(addr, srv)
}
