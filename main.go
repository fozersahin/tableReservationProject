package main

import (
	"fmt"
	"io"
	"net"
	"strings"
	"tableReservationProject/rest"
)

func main() {
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
		go handleRequest(conn)
	}
}

//TODO connection final check&timeout should be implemented
//TODO routing functionality should be implemented (ie:awareness of TCP to HTTP :  method:POST  route:/  protocol: HTTP/1.1 )

func handleRequest(conn net.Conn) {
	// Buffer that holds incoming information
	buf := make([]byte, 1024)
	length, err := conn.Read(buf)
	fmt.Println("Read len:", length)
	if length > 0 {

		request := string(buf[:length])
		var method = strings.TrimSpace(request[0:strings.Index(request, "/")])

		var body = request[strings.Index(request, "{"):len(request)]
		fmt.Print(body)
		if io.EOF == err {
			fmt.Println("EOF:")
		}

		rest.ReserveTable(body, method, conn)

	} else {
		conn.Write([]byte("HTTP/1.1 200 OK\r\n"))
		conn.Write([]byte("\r\n"))
		conn.Write([]byte("Request received!"))
		conn.Close()
	}
}
