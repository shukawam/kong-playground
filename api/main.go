package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	ua := ""
	router.Use(func(ctx *gin.Context) {
		ua = ctx.GetHeader("User-Agent")
		ctx.Next()
	})
	router.
		GET("/ping", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "pong")
		}).
		GET("/greet", func(ctx *gin.Context) {
			name := ctx.Query("name")
			ctx.JSON(http.StatusOK, gin.H{
				"message":    fmt.Sprintf("Hello world %s!!", name),
				"User-Agent": ua,
			})
		})
	router.Run("0.0.0.0:8080")
}
