package main

import (
	"fmt"
	"net/http"
	"time"
)

func GetHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		t := time.Now()

		// Format waktu sesuai permintaan
		formattedTime := fmt.Sprintf("%s, %d %s %d", t.Weekday().String(), t.Day(), t.Month().String(), t.Year())

		// Menampilkan waktu ke response writer
		writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
		writer.Write([]byte(formattedTime))

	} // TODO: replace this
}

func main() {
	http.ListenAndServe("localhost:8080", GetHandler())
}
