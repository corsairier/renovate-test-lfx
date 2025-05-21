package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

type client struct {
	conn *websocket.Conn
	send chan []byte
}

type broadcastMsg struct {
	sender *client
	data   []byte
}

var (
	clients   = make(map[*client]bool)
	broadcast = make(chan broadcastMsg)
	lock      = sync.Mutex{}
)

func handleConnections(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Upgrade failed: %v", err)
		return
	}
	c := &client{conn: ws, send: make(chan []byte, 256)}

	lock.Lock()
	clients[c] = true
	lock.Unlock()
	log.Printf("New Client conecected: %s", ws.RemoteAddr())

	go writeMessages(c)

	defer func() {
		lock.Lock()
		delete(clients, c)
		lock.Unlock()
		c.conn.Close()
	}()

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Printf("Read error: %v", err)
			break
		}
		broadcast <- broadcastMsg{sender: c, data: msg}
	}
}

func handleMessages() {
	for {
		b := <-broadcast
		lock.Lock()
		for c := range clients {
			if c == b.sender {
				continue
			}
			select {
			case c.send <- b.data:
			default:
				close(c.send)
				delete(clients, c)
			}
		}
		lock.Unlock()
	}
}

func writeMessages(c *client) {
	for msg := range c.send {
		err := c.conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Printf("Write error: %v", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleConnections)
	go handleMessages()

	fmt.Println("Chat server started at ws://localhost:8080/ws")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
