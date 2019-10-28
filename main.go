package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strconv"
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
		fmt.Println(conn)

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	message, err := bufio.NewReader(conn).ReadString('\n')

	fmt.Println(message)
	message = strings.Trim(message, " ")
	tableNum, err := strconv.Atoi(message)

	// Send table num, to func
	response := rest.ReserveTable(tableNum)

	if err != nil && err != io.EOF {
		fmt.Println("Error reading:", err.Error())
	}
	// Send a response back to person contacting us.
	conn.Write([]byte(response))
	// Close the connection when you're done with it.
	conn.Close()
}
