package main

import (
	"fmt"

	"github.com/muhwyndhamhp/gotes-mx/db"
	"github.com/muhwyndhamhp/gotes-mx/internal"
	"gorm.io/gorm"
)

func main() {
	db := db.GetDB()

	seedTodo(db)
}

func seedTodo(db *gorm.DB) {
	for i := 0; i < 10; i++ {
		if err := db.Create(&internal.Todo{
			Description: fmt.Sprintf("Todo Number %d", i),
			IsFinished:  false,
		}).Error; err != nil {
			panic(err)
		}
	}
}
