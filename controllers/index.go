package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
}

type ControllerError struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

func resp(c *gin.Context, data interface{}, newToken string) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
		"newToken": newToken,
	})
}

func fail(c *gin.Context, errs *ControllerError, newToken string) {
	c.JSON(http.StatusOK, gin.H{
		"code": errs.Code,
		"msg":  errs.Message,
		"newToken": newToken,
	})
}