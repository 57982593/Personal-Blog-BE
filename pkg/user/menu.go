package user

import (
	"encoding/json"
	database "goTestProject/init"
	"net/http"
)

type Menu struct {
	Title string
	Id int64
}
var menu []Menu

func (Menu) TableName() string {
	return "menu"
}

func GetMenuList(w http.ResponseWriter,r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Custom-Header", "custom")
	w.WriteHeader(201)
	db := database.GetDb()
	json,_ := json.Marshal(db.Find(&menu))
	w.Write(json)
}
