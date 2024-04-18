package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	users := readUsersFromFile("users.json")
	sort.Slice(users, func(i, j int) bool { return users[i].ID < users[j].ID })

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(users)
	})

	log.Println("Server listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func readUsersFromFile(filepath string) []User {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatalf("Error reading from file: %s", err)
	}
	var users []User
	json.Unmarshal(file, &users)
	return users
}
