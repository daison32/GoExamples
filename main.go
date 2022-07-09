package main

import (
	"github.com/gin-gonic/gin"
)

// サーバーを立ち上げる

func main() {
	r := gin.Default()
	r.GET("/tasks", tasksGet)
	r.Run()
}

