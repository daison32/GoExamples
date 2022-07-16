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
// GET
func (h handler) tasksGet(c *gin.Context) {

	tasks := []Task{}
	err := h.db1.Select(&tasks, "SELECT * FROM tasks")
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

// POST
func (h handler) tasksPost(c *gin.Context) {
	var postedItem postedStruct
	if err := c.ShouldBindJSON(&postedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err1 := h.db1.NamedExec(`INSERT INTO tasks (content) VALUES (:content)`,
		postedItem)
	if err1 != nil {
		log.Panic(err1)
	}

	c.JSON(200, gin.H{"string": postedItem.Content})
}

type postedStruct struct {
	Content string `db:"content" json:"content"`
}

// DELETE
func (h handler) tasksComplete(c *gin.Context) {
	var deletedItem deletedStruct
	if err := c.ShouldBindJSON(&deletedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err1 := h.db1.NamedExec(`DELETE FROM tasks WHERE id = (:id)`,
	deletedItem)
	if err1 != nil {
		log.Panic(err1)
	}

	c.JSON(200, gin.H{"string": deletedItem.ID})
}

type deletedStruct struct {
	ID int `db:"id" json:"id"`
}
