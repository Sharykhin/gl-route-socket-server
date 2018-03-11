package handler

import (
	"net/http"

	"github.com/Sharykhin/gl-route-socket-server/middleware"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
} // use default options

var connections = []*websocket.Conn{}

// Router provides new mux router for socket server
func Router() http.Handler {
	r := mux.NewRouter()
	r.Handle("/", middleware.Chain(http.HandlerFunc(echo), middleware.JWT))
	return r
}
