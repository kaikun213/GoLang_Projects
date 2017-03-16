package main

import (
	"log"
	"net"

	"github.com/kaikun213/go_experimential/networking/server/read"
)

func main() {
	// Server listens to incoming request
	ln, err := net.Listen("tcp", ":4950")
	if err != nil {
		log.Printf("Error listening to PORT 4950: %v", err)
	}
	// While loop
	for {
		// Accepts request
		conn, err := ln.Accept()
		if err != nil {
			panic(conn)
		}
		// handles request
		go handleConn(conn)
	}
}

// handle request
func handleConn(conn net.Conn) {
	// read file and format text
	var formattedText = read.Textfile("/home/kaikun/deinstalled.txt")
	conn.Write([]byte(formattedText))
}
