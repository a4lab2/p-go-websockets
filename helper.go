package main

import (
	"log"

	"github.com/gorilla/websocket"
)

func (app *application) reader(conn *websocket.Conn) {
	for {
		m, p, err := conn.ReadMessage()

		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(p))
		err = conn.WriteMessage(m, p)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
