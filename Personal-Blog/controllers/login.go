package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"goTestProject/Personal-Blog/data"
	"goTestProject/Personal-Blog/log"
	"io/ioutil"
)


func getReqBody(ctx *gin.Context) *data.User {
	buf := make([]byte, 1024)
	n, _ := ctx.Request.Body.Read(buf)
	ctx.Request.Body = ioutil.NopCloser(bytes.NewReader(buf[:n]))
	j := buf[0:n]
	stu := &data.User{}
	err := json.Unmarshal(j, &stu)
	if err != nil {
		panic(err)
	}
	return stu
}

func (c Controller) Login(ctx *gin.Context) {
	body := getReqBody(ctx)
	user, err := data.GetUser(body)
	if err != nil {
		log.Error(fmt.Sprint(err))
	}
	fmt.Println("user:", *user)
}

func (c Controller) Register(ctx *gin.Context) {
	body := getReqBody(ctx)
	err := data.Create(body)
	if err != nil {
		resp(ctx, err, "")
	}
	resp(ctx, "ok", "")
}
