package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// サーバーを立ち上げる

func main() {
	// DBに接続
	// db, err := sqlx.Open("postgres",
	// 	"user=postgres password=postgres dbname=postgres sslmode=disable")
	// if err != nil {
	// 	log.Panic(err)

	// }

	// retrieve the url
	dns := os.Getenv("DATABASE_URL")
	fmt.Println(dns)
	// connect to the db
	db, err := sqlx.Connect("postgres", dns)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	h := Newhandler(db)
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

	// トップページ（status:ok　と表示）
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// 一覧みる
	r.GET("/tasks", h.tasksGet)

	// とうろく
	r.POST("/tasks", h.tasksPost)

	r.Run()
}
