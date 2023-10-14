package sigma

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type (
	Router interface {
		Endpoint(method, pattern string, handler http.HandlerFunc)
		Use(middlewares ...func(next http.Handler) http.Handler)
		Group(pattern string, fn func(r Router))
		Static(pattern, path string)
		NotFound(handler http.HandlerFunc)
		NotAllowed(handler http.HandlerFunc)
		ServeHTTP(w http.ResponseWriter, r *http.Request)
	}

	base struct {
		router     *chi.Mux
		notFound   http.HandlerFunc
		notAllowed http.HandlerFunc
	}
)

func (s *Server) NewRouter() Router {
	return newBase(s.mux)
}

func newBase(r *chi.Mux) *base {
	return &base{
		router:     r,
		notFound:   http.NotFound,
		notAllowed: http.NotFound,
	}
}

func (b *base) Endpoint(method, pattern string, handler http.HandlerFunc) {
	b.router.Method(method, pattern, handler)
}

func (b *base) Use(mws ...func(next http.Handler) http.Handler) {
	for _, mw := range mws {
		b.router.Use(mw)
	}
}

func (b *base) Group(pattern string, fn func(r Router)) {
	r := newBase(chi.NewRouter())

	r.NotFound(b.notFound)
	r.NotAllowed(b.notAllowed)

	fn(r)

	b.router.Mount(pattern, r)
}

func (b *base) NotFound(handler http.HandlerFunc) {
	b.notFound = handler
	b.router.NotFound(handler)
}

func (b *base) NotAllowed(handler http.HandlerFunc) {
	b.notAllowed = handler
	b.router.MethodNotAllowed(handler)
}

func (b *base) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	b.router.ServeHTTP(w, r)
}

func (b *base) Static(pattern, path string) {
	b.router.Handle(pattern, http.StripPrefix(pattern, http.FileServer(http.Dir(path))))
}

func Param(r *http.Request, key string) string {
	return chi.URLParam(r, key)
}
