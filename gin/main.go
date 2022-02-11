package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//func server()  {
//
//
//}

func main() {
	r := gin.Default()
	r.Static("/static","./tem/")
	r.LoadHTMLGlob("tem/*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "我是测试", "ce": "123456"})
	})
	r.GET("/status", func(c *gin.Context) {
		c.HTML(http.StatusOK, "status.html", gin.H{"Num": "1","Host": "Host_H"})
	})
	r.POST("/post_context", func(c *gin.Context) {
		host_context := c.PostForm("host_client_context")
		host_status := c.PostForm("host_client_status")
		fmt.Println(host_context,host_status)
		c.String(http.StatusOK,host_status)

	})
	r.Run(":3000")
}