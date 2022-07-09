package main

import (
	"github.com/gin-gonic/gin"
)

// リクエストとレスポンスを書いていく

func tasksGet(c *gin.Context) {

	response := GetResponse{
		Tasks: []Task{
			{
				ID:         1,
				Content:    "しゅくだい",
				IsComplete: true,
			},
			{
				ID:         2,
				Content:    "おつかい",
				IsComplete: true,
			},
			{
				ID:         3,
				Content:    "さんぽ",
				IsComplete: false,
			},
		
		},
	}

	c.JSON(200, response)
}

type Task struct {
	ID         int    `json:"id"`
	Content    string `json:"content"`
	IsComplete bool   `json:"isComplete"`
}

type GetResponse struct {
	Tasks []Task `json:"tasks"`
}
