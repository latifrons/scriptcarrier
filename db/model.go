package db

import "gorm.io/gorm"

type Task struct {
	Name            string `gorm:"index"`
	ScriptType      string
	ScriptPath      string
	Args            string
	IntervalSeconds int
	gorm.Model
}
