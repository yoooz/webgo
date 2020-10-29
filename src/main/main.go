package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.Static("/static", "./static")
	ua := ""
	
	// ミドルウェア
	engine.Use(func(c *gin.Context) {
		ua = c.GetHeader("User-Agent")
		c.Next()
	})
	// html のディレクトリを指定
	engine.LoadHTMLGlob("templates/*")
	engine.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			// html に渡す変数を定義
			"message": "hello gin",
		})
	})
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message" : "ping",
			"User-Agent": ua,
		})
	})
	engine.Run(":3000")
	
}
