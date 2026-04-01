package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/morse", morseHandler)
	fmt.Println("Server starting on :8080")
	fmt.Println("Usage: GET /morse?text=HELLO")
	http.ListenAndServe(":8080", nil)
}
