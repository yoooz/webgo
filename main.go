package main

import (
	"webgo/controller"
	"webgo/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.Static("/static", "./static")
	// ミドルウェア
	engine.Use(middleware.RecordUaAndTime)
	// CRUD 書籍
	bookEngine := engine.Group("/book")
	{
		v1 := bookEngine.Group("/v1")
		{
			v1.POST("/add", controller.BookAdd)
			v1.GET("/list", controller.BookList)
			v1.PUT("/update", controller.BookUpdate)
			v1.DELETE("/delete", controller.BookDelete)
		}
	}

	engine.Run(":3000")

}
