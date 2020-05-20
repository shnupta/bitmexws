package bitmexws

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/shnupta/bitmexws/topic"
)

const (
	baseURI = "wss://www.bitmex.com/realtime"
)

var (
	Topics = topic.GetTopics()
)

// Opens a connection to the BitMEX WebSocket
func Connect() *websocket.Conn {
	conn, _, err := websocket.DefaultDialer.Dial(baseURI, nil)
	if err != nil {
		log.Fatalln(err)
	}
	return conn
}

// Subscribes to a topic
func Subscribe(send chan *Command, topic string) {
	cmd := CreateCommand("subscribe", topic)
	send <- cmd
}
