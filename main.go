package main

import (
	"database/sql"
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
	db, err := sqlx.Open("postgres",
		"user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		log.Panic(err)

	}

	// // retrieve the url
	// dbURL := os.Getenv("DATABASE_URL")
	// // connect to the db
	// db, err := sql.Open("postgres", connStr)

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

	r.GET("/tasks", h.tasksGet)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	r.Run()
}
