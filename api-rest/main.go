package main //package al que pertenece

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()

	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)
}
