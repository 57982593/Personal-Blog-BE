package main

import (
	"goTestProject/pkg/user"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/write", user.Write)
	err := http.ListenAndServe(":8080", nil)
	if (err) != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
