package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func welcome(c *gin.Context) {
	firstName := c.DefaultQuery("first_name", "ChuLiang")
	lastName := c.DefaultQuery("last_name", "Dong")

	example := c.MustGet("example")
	log.Println("---------------------example:", example)

	c.JSON(http.StatusOK, gin.H{
		"firstName": firstName,
		"lastName":  lastName,
	})
}

func formPost(c *gin.Context) {
	message := c.PostForm("message")
	nickName := c.DefaultPostForm("nick_name", "anonymous")

	c.JSON(http.StatusOK, gin.H{
		"status":    "posted",
		"message":   message,
		"nick_name": nickName,
	})

}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// 设置 example 变量
		c.Set("example", "12345")

		// 请求前

		c.Next()

		// 请求后
		latency := time.Since(t)
		log.Print(latency)

		// 获取发送的 status
		status := c.Writer.Status()
		log.Println(status)
	}
}

func main() {
	router := gin.Default()

	router.Use(Logger())
	router.GET("/welcome", welcome)
	router.POST("/form_post", formPost)

	router.Run(":8888")

}
