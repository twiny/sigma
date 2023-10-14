package sigma

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type (
	Server struct {
		mux *chi.Mux
		srv *http.Server
	}
)

func NewServer(addr string) *Server {
	mux := chi.NewRouter()

	return &Server{
		mux: mux,
		srv: &http.Server{
			Addr:         addr,
			Handler:      mux,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 20 * time.Second,
			IdleTimeout:  10 * time.Second,
		},
	}
}

func (s *Server) Start() error {
	log.Println("starting http server on", s.srv.Addr)
	return s.srv.ListenAndServe()
}

func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.srv.Shutdown(ctx); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}
