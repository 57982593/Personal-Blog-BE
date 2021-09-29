package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func (c Controller) DownloadMainExe(ctx *gin.Context) {
	filePath := "./main.exe"
	_, errByOpenFile := os.Open(filePath)
	//非空处理
	if errByOpenFile != nil {
		ctx.JSON(http.StatusOK, gin.H{
		    "success": false,
		    "message": "失败",
		    "error":   "资源不存在",
		})
		fmt.Println(errByOpenFile)
		ctx.Redirect(http.StatusFound, "/404")
		return
	}
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Content-Disposition", "attachment; filename=main.exe")
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.File(filePath)
	return
}
