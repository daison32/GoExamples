package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// 構造体に、DBを入れる
type handler struct {
	db1 *sqlx.DB
}

// ↑のhandlerのコンストラクター（初期化するための関数）
func Newhandler(d *sqlx.DB) handler {
	return handler{
		db1: d,
	}
}

// リクエストとレスポンスを書いていく
func (h handler) tasksGet(c *gin.Context) {

	tasks := []Task{}
	err := h.db1.Select(&tasks, "select * from tasks")
	if err != nil {
		log.Panic(err)
	}

	response := GetResponse{
		Tasks: tasks,
	}

	c.JSON(200, response)
}

type Task struct {
	ID         int    `db:"id" json:"id"`
	Content    string `db:"content" json:"content"`
	IsComplete bool   `db:"is_completed" json:"isComplete"`
}

type GetResponse struct {
	Tasks []Task `json:"tasks"`
}
