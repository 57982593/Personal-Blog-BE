package user

import (
	json2 "encoding/json"
	"fmt"
	database "goTestProject/init"
	"net/http"
)

type Router struct {
	Route string
	Id int64
}

func (Router)TableName() string {
	return "router"
}

func GetRoute(w http.ResponseWriter,r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Custom-Header", "custom")
	w.WriteHeader(201)
	db := database.GetDb()
	//data := r.URL.Query()
	//fmt.Println(data.Get("id"), db)
	var routers []Router
	route := db.Find(&routers)
	fmt.Printf("%#v",route)
	json,_:= json2.Marshal(route)
	w.Write(json)
}
