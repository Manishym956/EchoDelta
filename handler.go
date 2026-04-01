package main

import (
	"net/http"
	"strings"
)

func morseHandler(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("text")
	if text == "" {
		http.Error(w, "Missing 'text' query parameter", http.StatusBadRequest)
		return
	}
	morse := textToMorse(strings.ToUpper(text))
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(morse))
}
