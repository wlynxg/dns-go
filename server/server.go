package server

import (
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
}

func (s *Server) Start() error {
	return s.server.ListenAndServe()
}

func (s *Server) Close() error {
	return s.server.Close()
}

func New() *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/update", Update)
	server := &http.Server{
		Addr:         ":8081",
		WriteTimeout: 10 * time.Second,
		Handler:      mux,
	}

	return &Server{
		server: server,
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
