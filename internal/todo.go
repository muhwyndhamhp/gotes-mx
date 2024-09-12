package internal

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model

	Description string
	IsFinished  bool
	FinishedAt  *time.Time
}
