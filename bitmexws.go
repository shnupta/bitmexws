// A wrapper for the BitMEX WebSocket API.
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
// Commands will be sent through the `send` channel.
func Subscribe(send chan *Command, topic string) {
	cmd := CreateCommand("subscribe", topic)
	send <- cmd
}

// TODO:
// Decide on ways that you can interact with channels. i.e. Should you be able
// to assign different channels to different response messages?
// So if a "trade" object is sent then it will be routed to a specified channel?
// No I think this should be left to the user to do.
//
// On another point, I don't think the read channel should send through raw
// []byte. I think I should handle this and return a standard struct that the
// user can open and get one of the three response structs.
