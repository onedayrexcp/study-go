package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)
import "net/http"

func MiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		now := time.Now()
		fmt.Println("smart req start ==>")
		since := time.Since(now)
		fmt.Println("time ==> ", since)
	}
}

func main() {
	r := gin.Default()
	r.Use(MiddleWare())
	r.GET("/user/:name/*action", func(context *gin.Context) {
		name := context.Param("name")
		action := context.Param("action")
		context.JSON(http.StatusOK, "Hello World==>"+name+action)
	})
	r.GET("/param", func(context *gin.Context) {
		name := context.DefaultQuery("name", "default")
		context.JSON(http.StatusOK, "Hello World==>"+name)
	})
	r.Run(":8888")
}
