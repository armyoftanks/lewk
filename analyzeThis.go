package main

import (
	"fmt"
	"log"

	"golang.org/x/net/websocket"
)

// create twilio 2 way messenger
// create a websocket
// implement microsoft vision api
// feed image data back into twilio through websocket

func Dial() {
	// do i need this?
}

type data struct {
}

func webSock() {

	//attempting to open a websocket connection
	origin := "http://localhost/"
	url := "ws://localhost:12345/ws"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := ws.Write([]byte("hello, world!\n")); err != nil {
		log.Fatal(err)
	}
	var msg = make([]byte, 512)
	var n int
	if n, err = ws.Read(msg); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Received: %s.\n", msg[:n])

}

func chat() {

}

func analyzer() {

}

func main() {

}
