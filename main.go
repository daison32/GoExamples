package main

import (
	"/handler.go"

	"github.com/gin-gonic/gin"
)
// func main() {
// 	r := gin.Default()
// 	r.GET("/health", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"status": "ok",
// 		})
// 	})
// 	// r.Run(":80") ｚ
// }

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
        // 値を取得する
		tasks := handler.GetAll()
        // JSONで返す
		c.JSON(200, tasks)
	})

	r.Run()
}