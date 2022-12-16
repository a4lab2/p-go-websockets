package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func (app *application) homepage(w http.ResponseWriter, r *http.Request) {

}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (app *application) websocket(w http.ResponseWriter, r *http.Request) {

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)

	}
	log.Println("Client Successfully Connected")
	app.reader(ws)

}
