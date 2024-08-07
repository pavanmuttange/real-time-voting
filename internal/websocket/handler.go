package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	conn *websocket.Conn
	send chan Message
}

func (c *Client) readPump() {
	defer func() {
		c.conn.Close()
	}()
	for {
		var msg Message
		err := c.conn.ReadJSON(&msg)
		if err != nil {
			log.Println("Error reading JSON:", err)
			break
		}

		session := GetSession(msg.Session)
		if session == nil {
			session = CreateSession(msg.Session)
		}
		session.votes[msg.Vote]++
		session.Broadcast(msg)
	}
}

func (c *Client) WritePump() {
	defer func() {
		c.conn.Close()
	}()
	for {
		msg := <-c.send
		err := c.conn.WriteJSON(msg)
		if err != nil {
			log.Println("error writing JSON:", err)
			break
		}

	}
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("error in upgrading connection:", err)
		return
	}
	client := &Client{
		conn: conn,
		send: make(chan Message),
	}

	sessionID := r.URL.Query().Get("session")
	session := GetSession(sessionID)
	if session == nil {
		session = CreateSession(sessionID)
	}
	session.AddClient(client)
	defer session.RemoveClient(client)

	go client.WritePump()
	client.readPump()
}
