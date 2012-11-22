package main

import (
	// go get -u code.google.com/p/go.net
	"code.google.com/p/go.net/websocket"
	"fmt"
	"net/http"
)

type connection struct {
	// The websocket connection.
	ws *websocket.Conn

	// Buffered channel of outbound messages.
	send chan string
}

func (c *connection) reader() {
	for {
		var message string
		err := websocket.Message.Receive(c.ws, &message)
		c.send <- message
		fmt.Println("read", message)
		if err != nil || message == "close" {
			break
		}
	}
	c.ws.Close()
}

func (c *connection) writer() {
	for message := range c.send {
		err := websocket.Message.Send(c.ws, message)
		if err != nil {
			break
		}
	}
	c.ws.Close()
}

func wsHandler(ws *websocket.Conn) {

	c := &connection{send: make(chan string, 256), ws: ws}
	go c.writer()
	c.reader()


	// websocket.Message.Send(ws, "this is a test")
	// var message string
	// websocket.Message.Receive(ws, &message)
	// fmt.Println("received", message)
	// ws.Close()
}

func main() {
	http.Handle("/s/", http.StripPrefix("/s/", http.FileServer(http.Dir("s"))))
	http.Handle("/ws", websocket.Handler(wsHandler))
	// fmt.Println("abc"== "aaa")
	err := http.ListenAndServe(":9898", nil)
	fmt.Println(err)
}
