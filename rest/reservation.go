package rest

import (
	"encoding/json"
	"fmt"
	"net"
)

type reservation struct {
	Table int
}

var reservations = make(map[int]reservation)

//TODO Simplify connection handling
func ReserveTable(body string, method string, conn net.Conn) {
	if method == "POST" {
		var newReservation reservation
		json.Unmarshal([]byte(body), &newReservation)

		if newReservation.Table <= 0 || newReservation.Table > 30 {

			conn.Write([]byte("HTTP/1.1 400 Bad Request \r\n"))
			conn.Write([]byte("\r\n"))
			conn.Write([]byte("Enter a valid number!"))
			conn.Close()
		}

		if value, ok := reservations[newReservation.Table]; ok {
			fmt.Println("ERROR- Table reserved: ", value.Table)
			conn.Write([]byte("HTTP/1.1 400 Bad Request\r\n"))
			conn.Write([]byte("\r\n"))
			conn.Write([]byte("Table has already been reserved!"))
			conn.Close()

		} else {
			fmt.Println("Table reserving.. Table no:", newReservation.Table)
			reservations[newReservation.Table] = newReservation
			conn.Write([]byte("HTTP/1.1 201 Created \r\n"))
			conn.Write([]byte("\r\n"))
			conn.Write([]byte("Table reserved for you!"))
			conn.Close()
		}

	} else {
		conn.Write([]byte("HTTP/1.1 405 Method Not Allowed \r\n"))
		conn.Write([]byte("\r\n"))
		conn.Close()

	}
}
