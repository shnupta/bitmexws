package main

import (
	"log"
	"os"

	"github.com/gorilla/websocket"
	"github.com/shnupta/bitmexws"
)

func readMessages(conn *websocket.Conn, read chan []byte) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Fatalln(err)
		}

		read <- message
	}
}

func sendMessages(conn *websocket.Conn, send chan *bitmexws.Command) {
	for {
		cmd := <-send
		conn.WriteJSON(cmd)
	}
}

func main() {
	ws := bitmexws.Connect()
	defer ws.Close()

	send := make(chan *bitmexws.Command, 256)
	read := make(chan []byte, 256)

	go readMessages(ws, read)
	go sendMessages(ws, send)

	//bitmexws.Subscribe(send, bitmexws.Topics.Connected)
	apiKey := os.Getenv("BITMEX_KEY")
	apiSecret := os.Getenv("BITMEX_SECRET")
	bitmexws.Authenticate(send, apiKey, apiSecret)

	for {
		msg := <-read
		log.Println(string(msg))
	}
}
