package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/twiny/sigma"
	"github.com/twiny/sigma/middleware"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	srv := sigma.NewServer(":81")
	defer srv.Stop()

	router := srv.NewRouter()

	router.Use(
		middleware.CORS([]string{"*"}),
		middleware.RealIP,
		middleware.Gzip(5),
	)

	router.Endpoint(http.MethodGet, "/event/{name}", echo)

	srv.Start()
}

func echo(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	//
	tik := time.NewTicker(time.Second)
	for {
		select {
		case <-tik.C:
			if err := c.WriteMessage(websocket.TextMessage, []byte("hello")); err != nil {
				log.Println("write:", err)
				return
			}
		}
	}
}
