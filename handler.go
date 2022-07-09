package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	// db_model.goでは使わないが，内部的に使われるためここでimportしておく
	"github.com/mattn/go-sqlite3"
)

type tasks struct {
	// 大文字にしなければ値が取得できない
	id   int    `gorm:"primary_key"`
	text string `gorm:"not null"`
}

func GetAll() []tasks {
	db, err := gorm.Open("sqlite3", "data/taskss.db")
	if err != nil {
		panic("failed to connect database")
	}

	taskss := []tasks{}

	// 全ての値を取得
	db.Find(&taskss)

	return taskss
}

func GetBy(tag string, num string) []tasks {
	db, err := gorm.Open("sqlite3", "data/taskss.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	taskss := []tasks{}
	db.Where(fmt.Sprintf("%s = ?", tag), num).Find(&taskss)

	return taskss
}
