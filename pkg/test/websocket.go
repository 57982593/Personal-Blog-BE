package test

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
func Websocket(w http.ResponseWriter, r * http.Request) {
	var (
		wbsCon *websocket.Conn
		err error
		data []byte
	)
	fmt.Println(r.URL.Query())
	// 完成http应答，在httpheader中放下如下参数
	if wbsCon, err = upgrader.Upgrade(w, r, nil);err != nil {
		fmt.Printf("err:--%v\n", err)
		return // 获取连接失败直接返回
	}

	for {
		// 只能发送Text, Binary 类型的数据,下划线意思是忽略这个变量.
		if _, data, err = wbsCon.ReadMessage();err != nil {
			fmt.Printf("err:--%v\n", err)
			goto ERR // 跳转到关闭连接
		}
		if err = wbsCon.WriteMessage(websocket.TextMessage, data); err != nil {
			fmt.Printf("err:--%v\n", err)
			goto ERR // 发送消息失败，关闭连接
		}
	}

ERR:
	// 关闭连接
	wbsCon.Close()
}
