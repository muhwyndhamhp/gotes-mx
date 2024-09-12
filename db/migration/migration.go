package main

import (
	"github.com/muhwyndhamhp/gotes-mx/db"
	"github.com/muhwyndhamhp/gotes-mx/internal"
)

func main() {
	db := db.GetDB()

	if err := db.AutoMigrate(&internal.Todo{}); err != nil {
		panic(err)
	}

	db.Debug()
}
