package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getting(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "someGet",
	})
}

type Person struct {
	ID   string `uri:"id" binding:"required"`
	Name string `uri:"name" binding:"required"`
}

func main() {

	router := gin.Default()

	goods := router.Group("/goods")
	{
		// 共同前缀/goods
		goods.GET("/lists", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "goods/lists",
			})
		})
	}

	router.GET("/:name/:id", func(c *gin.Context) {
		var p Person
		err := c.ShouldBindUri(&p)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": err,
			})
			return
		}
		fmt.Println(p)
		c.JSON(200, gin.H{"name": p.Name, "id": p.ID})
	})

	// 默认启动的是 8080端口，也可以自己定义启动端口
	router.Run(":8888")
}
