# sigma

a small wrapper around [go-chi](https://github.com/go-chi/chi) HTTP router.


**NOTE**: This package is provided "as is" with no guarantee. Use it at your own risk and always test it yourself before using it in a production environment. If you find any issues, please [create a new issue](https://github.com/twiny/sigma/issues/new).

## API
Router methods.
```go
type Router interface {
	Endpoint(method, pattern string, handler http.HandlerFunc)
	Use(middlewares ...func(next http.Handler) http.Handler)
	Group(pattern string, fn func(r Router))
	NotFound(handler http.HandlerFunc)
	NotAllowed(handler http.HandlerFunc)
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}
```
Param returns the url parameter from a http.Request object.
```go
Param(r *http.Request, key string) string
```

## Example
```go
package main

import (
	"log"
	"net/http"

	"github.com/twiny/sigma"
)

func main() {
	srv := sigma.NewServer(":1234")
	defer srv.Stop()

	router := srv.NewRouter()

	// custom 404 not found handler
	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Not Found"))
	})

	// custom 405 not allowed handler
	router.NotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 - Method Not Allowed"))
	})

	// middleware
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r.Header.Add("x-key-1", "main router")
			next.ServeHTTP(w, r)
		})
	})

	// endpoints
	router.Endpoint(http.MethodGet, "/hello/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := sigma.Param(r, "name")
		w.Write([]byte("Hello World " + name))
	})

	// sub router
	router.Group("/v1", func(r sigma.Router) {
		// sub router middleware
		r.Use(func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				r.Header.Add("x-key-2", "sub router")
				next.ServeHTTP(w, r)
			})
		})
		//
		r.Endpoint(http.MethodGet, "/hello/{name}", func(w http.ResponseWriter, r *http.Request) {
			name := sigma.Param(r, "name")
			w.Write([]byte("Hello World v1 " + name))
		})
	})

	log.Fatal(srv.Start())
}
```

