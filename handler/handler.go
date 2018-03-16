package handler

import (
	"fmt"
	"log"
	"net/http"

	"sync"

	"encoding/json"

	"github.com/Sharykhin/gl-route-socket-server/entity"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	m     = sync.Mutex{}
	users = []Connection{}
)

type Connection struct {
	conn *websocket.Conn
	mu   sync.Mutex
}

func handle(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)

	fmt.Println("new connection is coming...")

	if err != nil {
		fmt.Printf("could not upgrade connection to websocket: %v", err)
		return
	}
	defer c.Close()

	appendConnection(c)

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			if e, ok := err.(*websocket.CloseError); ok {
				removeCon(c, e)
			}
			fmt.Printf("could not read a message: %v", err)
			break
		}

		log.Printf("recv: %s, type %v", message, mt)

		var em entity.Message
		err = json.Unmarshal(message, &em)
		if err != nil {
			fmt.Printf("could not parse income message: %v", err)
			break
		}

		switch em.Action {
		case "message":
			sendMessage(c, mt, em)
		}

	}
}

func sendMessage(c *websocket.Conn, mt int, em entity.Message) {
	var payload entity.MessagePayload
	err := json.Unmarshal(em.Payload, &payload)
	if err != nil {
		fmt.Printf("could not parse payload: %v", err)
		return
	}
	for _, u := range users {
		if u.conn != c {
			fmt.Printf("send message from %s \n", payload.Text)
			u.mu.Lock()
			err := u.conn.WriteMessage(mt, []byte(payload.Text))
			u.mu.Unlock()
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	}
}

func removeCon(c *websocket.Conn, e *websocket.CloseError) {
	fmt.Printf("connection has been closed by client. Code: %v\n", e.Code)
	for i, u := range users {
		if u.conn == c {
			m.Lock()
			users = append(users[:i], users[i+1:]...)
			m.Unlock()
		}
	}
	fmt.Println(users)
}

func appendConnection(c *websocket.Conn) {
	connection := Connection{
		conn: c,
		mu:   sync.Mutex{},
	}

	m.Lock()
	users = append(users, connection)
	m.Unlock()
	fmt.Println(users)
}
