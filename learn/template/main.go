package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"text/template"
)


func Dome() {
	fmt.Println("123")
	cmd := exec.Command("cmd.exe","/c", "node", "./javascript/pdf.js")
	e := cmd.Run()
	if e != nil {
		panic(e)
	}
}

func sayHello() {
	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("./learn/exampleTemplate/pdfTemplate.html")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	outFile := "./static/test.html"
	file, e2 := os.OpenFile(outFile, os.O_CREATE|os.O_WRONLY, 0755)
	if e2 != nil {
		panic(e2)
	}
	// 利用给定数据渲染模板，并将结果写入w
	e3 := tmpl.Execute(file, "5lmh.com")
	if e3 != nil {
		panic(e3)
	}
	// TODO 死循环尚未解决
	//go Dome()

}

func main() {
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	sayHello()
	defer Dome()

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("HTTP server failed,err:", err)
		return
	}

}
