package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// 構造体に、DBを入れるKK
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

// 登録


func (h handler) tasksPost(c *gin.Context) {
	var json Task
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"int": json.ID, "string": json.Content, "bool": json.IsComplete})

	_, err1 := h.db1.NamedExec(`INSERT INTO tasks (content) VALUES (:content)`, 
        map[string]interface{}{

            "content": "お試しPOST",
    })
	if err1 != nil {
		log.Panic(err1)
	}
}


type JsonRequest struct {
	Content    string `db:"content" json:"content"`
	IsComplete bool   `db:"is_completed" json:"isComplete"`
}
