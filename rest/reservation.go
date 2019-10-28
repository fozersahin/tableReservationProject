package rest

import (
	"fmt"
)

type reservation struct {
	Table int
}

var reservations = make(map[int]reservation)

func ReserveTable(tableNumber int) string {
	var newReservation reservation

	// for string requests , table number will be zero.
	newReservation.Table = tableNumber
	if newReservation.Table <= 0 || newReservation.Table > 30 {
		return "Invalid table number!\n"
	}

	// availability check
	if value, ok := reservations[newReservation.Table]; ok {
		fmt.Println("ERROR- Table is already reserved: ", value.Table)
		return "Table has already been reserved!\n"
	} else {
		fmt.Println("Table reserving.. Table no:", newReservation.Table)
		reservations[newReservation.Table] = newReservation
		return "Table is reserved for you\n"
	}
}
