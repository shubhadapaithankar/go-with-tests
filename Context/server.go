package main

import (
	"fmt"
	"net/http"
)

// Server function that takes a Store and returns an http.HandlerFunc.
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())
		if err != nil {
			// Handle the error, e.g., log it or return an HTTP error response.
			http.Error(w, "Could not fetch data", http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, data)
	}
}
