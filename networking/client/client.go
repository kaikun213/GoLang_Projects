package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "address")
	if err != nil {
		panic(conn)
	}

	fmt.Fprintf(conn, "GET FILE")

	data, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		panic(conn)
	}

	fmt.Printf("Content: %s", data)
}
