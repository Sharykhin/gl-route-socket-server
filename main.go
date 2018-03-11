package main

import (
	"log"
	"net/http"
	"os"

	"fmt"

	"github.com/Sharykhin/gl-route-socket-server/handler"
)

var address string

func init() {
	var address = os.Getenv("HTTP_ADDRESS")
	if address == "" {
		address = "127.0.0.1:1234"
	}
}

func main() {
	fmt.Printf("Listen on %s\n", address)
	log.Fatal(http.ListenAndServe(address, handler.Router()))
}
