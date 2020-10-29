package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"io"
	"log"
	"os"
)

func main() {
	engine := gin.Default()
	engine.Static("/static", "./static")
	
	// html のディレクトリを指定
	engine.LoadHTMLGlob("templates/*")
	engine.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			// html に渡す変数を定義
			"message": "hello gin",
		})
	})
	ping(engine)
	upload(engine)
	engine.Run(":3000")
	
}

func ping(engine *gin.Engine) {
	ua := ""
	// ミドルウェア
	engine.Use(func(c *gin.Context) {
		ua = c.GetHeader("User-Agent")
		c.Next()
	})
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message" : "ping",
			"User-Agent": ua,
		})
	})
}

func upload(engine *gin.Engine) {
	engine.POST("/upload", func(c *gin.Context) {
		file,header, err := c.Request.FormFile("image")
		if err != nil {
			c.String(http.StatusBadRequest, "Bad request")
			return
		}
		fileName := header.Filename
		dir, _ := os.Getwd()
		out, err := os.Create(dir+"\\images\\"+fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
}
