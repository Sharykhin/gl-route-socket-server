package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	connections = []*websocket.Conn{}
)

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	connections = append(connections, c)
	fmt.Println(connections)
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			if e, ok := err.(*websocket.CloseError); ok {
				fmt.Printf("connection has been closed by client. Code: %v\n", e.Code)
				for i, v := range connections {
					if v == c {
						connections = append(connections[:i], connections[i+1:]...)
					}
				}
			}
			log.Println("read:", err)

			break
		}
		log.Printf("recv: %s", message)

		go func() {
			for _, cc := range connections {
				if cc != c {
					err = cc.WriteMessage(mt, message)
					if err != nil {
						log.Println("write:", err)
						break
					}
				}
			}
			time.Sleep(15 * time.Second)
		}()
	}
}
