package api

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
	http   *http.Server
}

func NewServer(port int, uh *Handler) *Server {
	router := mux.NewRouter()
	router.HandleFunc("/redis/incr", uh.Incr)
	router.HandleFunc("/sign/hmacsha512", uh.GenerateHash)
	router.HandleFunc("/postgres/users", uh.AddUser)

	httpServer := &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: router,
	}
	server := &Server{
		router: router,
		http:   httpServer,
	}
	return server
}

func (s *Server) Run() error {
	return s.http.ListenAndServe()
}
