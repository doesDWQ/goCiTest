package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello kitty"))
	})

	log.Fatal(http.ListenAndServe(":8821",nil))
}
