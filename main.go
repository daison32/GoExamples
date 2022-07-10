package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// サーバーを立ち上げる

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		// アクセス許可するオリジン
		AllowOrigins: []string{
			"*",
		},
		// アクセス許可するHTTPメソッド
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPut,
			http.MethodPost,
			http.MethodDelete,
			http.MethodPatch,
		},
		// 許可するHTTPリクエストヘッダ
		AllowHeaders: []string{
			"*",
		},
	}))

	r.GET("/tasks", tasksGet)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	r.Run()
}
