package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type reservation struct {
	Table int
}

var reservations = make(map[int]reservation)

func ReserveTable(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var newReservation reservation
		decoder := json.NewDecoder(r.Body)
		decoder.Decode(&newReservation)

		if newReservation.Table <= 0 || newReservation.Table > 30 {
			w.WriteHeader(400)
			w.Write([]byte("Invalid table number!"))
			return
		}

		if value, ok := reservations[newReservation.Table]; ok {
			fmt.Println("ERROR- Table reserved: ", value.Table)
			w.WriteHeader(400)
			w.Write([]byte("Table has already been reserved!"))
		} else {
			fmt.Println("Table reserving.. Table no:", newReservation.Table)
			reservations[newReservation.Table] = newReservation
			w.WriteHeader(201)
			w.Write([]byte("Table is reserved for you"))
		}

	} else {
		http.Error(w, "Method not found.", http.StatusMethodNotAllowed)
	}
}
