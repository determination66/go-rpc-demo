package main

import (
	"github.com/gin-gonic/gin"
	"go_rpc_demo/demo3/protoc"
	"net/http"
)

type Login struct {
	User     string `form:"user" json:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main() {
	r := gin.Default()
	r.GET("/welcome", welcome)
	r.GET("/return_proto", returnProto)

	//会将特殊的html字符替换成unicode,例如<替换成\u003cb
	r.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"html": "<b>Hello World!</b>",
		})
	})

	// 原样输出
	r.GET("/pure_get", func(c *gin.Context) {
		c.PureJSON(http.StatusOK, gin.H{
			"html": "<b>Hello World!</b>",
		})
	})

	// login 传入json
	r.POST("/login_json", func(c *gin.Context) {
		var login Login
		err := c.ShouldBindJSON(&login)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err,
			})
			return
		}

		if login.User != "zhangsan" || login.Password != "nihao" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "StatusUnauthorized",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "login success!",
		})

	})

	// login 传入form
	r.POST("/login_form", func(c *gin.Context) {
		var login Login
		err := c.ShouldBind(&login)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err,
			})
			return
		}

		if login.User != "zhangsan" || login.Password != "nihao" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "StatusUnauthorized",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "login success!",
		})
	})

	r.Run(":8888")
}

func returnProto(c *gin.Context) {
	user := protoc.User{
		Name:    "姓名",
		Message: "信息",
		Number:  12,
		Books:   []string{"go", "gin", "微服务"},
	}

	c.ProtoBuf(http.StatusOK, &user)

}

func welcome(c *gin.Context) {
	var msg struct {
		Name    string
		Message string
		Number  int
	}

	msg.Name = "bobby"
	msg.Message = "这个是以测试json"
	msg.Number = 19

	c.JSON(http.StatusOK, msg)

}
