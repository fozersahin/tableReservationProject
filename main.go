package main

import (
	"net/http"
	"tableReservationProject/rest"
)

func main() {
	http.HandleFunc("/", rest.ReserveTable)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
