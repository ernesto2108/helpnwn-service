package account_service

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

type Server struct {
	*http.Server
}

func MServer(mux *chi.Mux) *Server {
	s := &http.Server{
		Addr:           ":9000",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &Server{s}

}

func (s *Server) Run() {
	log.Fatal(s.Server.ListenAndServe())
}
