package user

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Name   string
	Habits []string
}

func Write(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Custom-Header", "custom")
	w.WriteHeader(201)
	user := &User{
		Name:   "wang",
		Habits: []string{"balls", "running", "hiking"},
	}

	json, _ := json.Marshal(user)
	w.Write(json)
}
