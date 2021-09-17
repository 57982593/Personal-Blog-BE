package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"goTestProject/Personal-Blog/data"
	"io/ioutil"
)


func getReqBody(ctx *gin.Context) *data.User {
	buf := make([]byte, 1024)
	n, _ := ctx.Request.Body.Read(buf)
	ctx.Request.Body = ioutil.NopCloser(bytes.NewReader(buf[:n]))
	j := buf[0:n]
	stu := &data.User{}
	_ = json.Unmarshal(j, &stu)
	return stu
}

func (c Controller) Login(ctx *gin.Context) {
	body := getReqBody(ctx)
	user := data.GetUser(body)
	fmt.Println("user:", *user)
}

func (c Controller) Register(ctx *gin.Context) {
	body := getReqBody(ctx)
	data.Create(body)
	resp(ctx, "ok", "")
}
