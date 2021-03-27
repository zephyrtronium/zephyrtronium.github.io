package main

import (
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8008", http.FileServer(http.Dir("."))))
}
