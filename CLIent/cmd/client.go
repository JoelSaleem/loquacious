package main

import (
	"fmt"
	"net/url"
	"os"

	"golang.org/x/net/websocket"
)

func main() {
	// Parse the URL of the WebSocket server.
	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/connect"}

	// Connect to the server.
	ws, err := websocket.Dial(u.String(), "", u.String())
	if err != nil {
		fmt.Fprintf(os.Stderr, "dial: %v\n", err)
		os.Exit(1)
	}

	// Send a message to the server.
	if _, err := ws.Write([]byte("ping")); err != nil {
		fmt.Fprintf(os.Stderr, "write: %v\n", err)
		os.Exit(1)
	}

	// Read a message from the server.
	var msg = make([]byte, 512)
	n, err := ws.Read(msg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "read: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Received: %s\n", msg[:n])

	// Close the connection.
	if err := ws.Close(); err != nil {
		fmt.Fprintf(os.Stderr, "close: %v\n", err)
		os.Exit(1)
	}
}
