package main

import (
	"encoding/json"
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

	apiKey := os.Getenv("BITMEX_KEY")
	apiSecret := os.Getenv("BITMEX_SECRET")
	bitmexws.Authenticate(send, apiKey, apiSecret)
	bitmexws.Subscribe(send, bitmexws.Topics.Connected)

	for {
		msg := <-read
		var suc bitmexws.SubSuccess
		var serr bitmexws.SubError
		var data bitmexws.SubData

		if err := json.Unmarshal(msg, &suc); err != nil {
			log.Fatalln(err)
		}
		if !suc.Success {
			if err := json.Unmarshal(msg, &serr); err != nil {
				log.Fatalln(err)
			}
			if serr.Error == "" {
				if err := json.Unmarshal(msg, &data); err != nil {
					log.Fatalln(err)
				}
				if data.Table == "" {

				} else {
					log.Printf("Data subscription. Table: %s", data.Table)
				}
			} else {

				log.Printf("Error response: %s", serr.Error)
			}
		} else {
			log.Println(string(msg))
			log.Printf("Subscription success: %s", suc.Subscribe)
		}
	}
}
