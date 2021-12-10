package sigma

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Server http server
type Server struct {
	mux *chi.Mux
	srv *http.Server
}

// NewServer
func NewServer(addr string) *Server {
	mux := chi.NewRouter()

	mux.Use(
		middleware.Compress(5, "gzip"),
		middleware.Timeout(30*time.Second),
		middleware.RealIP,
		middleware.StripSlashes,
	)

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

// Start server
func (s *Server) Start() error {
	log.Println("starting http server on", s.srv.Addr)
	return s.srv.ListenAndServe()
}

// Stop server
func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.srv.Shutdown(ctx)
}
