package test

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
func CreateTable(w http.ResponseWriter,r *http.Request) {
	db := database.GetDb()
	fmt.Println(r.URL.Query())
	req := r.URL.Query()
	route := Router{Route: req.Get("route")}
	fmt.Println(db.NewRecord(&route))
	err := db.Create(&route)
	fmt.Println(db.NewRecord(&route))
	if err.Error!=nil {
		fmt.Println(err.Error)
		panic("err")
	}
	json,_:= json2.Marshal("插入成功")
	w.Write(json)
}
func SelectAll(w http.ResponseWriter,r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("X-Custom-Header", "custom")
	//w.WriteHeader(201)
	db := database.GetDb()
	var routers []Router
	data := db.Find(&routers)
	json,err := json2.Marshal(data)
	if err!=nil {
		fmt.Println("出错了")
	}
	w.Write(json)
}
