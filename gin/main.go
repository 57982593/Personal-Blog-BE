package main

import (
	"fmt"
	"goTestProject/gin/model"
	database "goTestProject/init"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	database.InitDatabase()
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST, GET, OPTIONS, PUT, DELETE, UPDATE"},
		AllowHeaders: []string{"Origin, X-Requested-With, Content-Type, Accept, Authorization"},
		ExposeHeaders: []string{"Content-Length, Access-Control-Allow-Origin, " +
			"Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))
	// This handler will match /user/john but will not match /user/ or /user
	// 这里可以使用 /user/自定义路由名称 来访问， c.Param 会接收到自定义得路由名称。
	router.GET("/getUserList", func(c *gin.Context) {
		users := model.GetUserList()
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"code":    200,
			"data":    users,
		})
	})

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	//多重自定义路由，效果同上。
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})
	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe
	// 在url后携带参数
	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest") //接受请求参数没有时设置默认值
		// 仅接受参数，但不设置默认值
		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	router.POST("/form_post", func(c *gin.Context) {
		//fmt.Printf("%+v\n", c)
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")
		fmt.Println(c.Query("message"))
		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	router.Run(":8083")
}
