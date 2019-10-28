package main

import (
	"fmt"
	"net"
)

func main() {
	//http.HandleFunc("/", rest.ReserveTable)
	//if err := http.ListenAndServe(":8080", nil); err != nil {
	//	panic(err)
	//}

	socket()
}

func socket() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		fmt.Println(conn)

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	message, err := conn.Read(buf)

	fmt.Println(message)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	// Send a response back to person contacting us.
	conn.Write([]byte("Message received.\n"))
	// Close the connection when you're done with it.
	conn.Close()
}
