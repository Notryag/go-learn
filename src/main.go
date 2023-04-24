package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello Workd")
	})

	r.GET("/user/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.String(http.StatusOK, "hello %s", name)
	})

	r.GET("/user/:name/*action", func(ctx *gin.Context) {
		name := ctx.Param("name")
		action := ctx.Param("action")
		ctx.JSON(http.StatusOK, gin.H{
			"name":   name,
			"action": action,
		})
	})
	// url参数
	r.GET("/user", func(ctx *gin.Context) {
		name := ctx.DefaultQuery("name", "guest")
		age := ctx.Query("age")
		ctx.String(http.StatusOK, "hello %s, age %s", name, age)
	})
	// post参数
	r.POST("/user", func(ctx *gin.Context) {
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")
		ctx.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/hello", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "hello",
			})
		})
	}
	r.Run(":8000")
}
