package db

import "gorm.io/gorm"

type Task struct {
	Name            string `gorm:"unique;size:128"`
	ScriptType      string `gorm:"size:64"`
	ScriptPath      string `gorm:"size:128"`
	Args            string `gorm:"size:1024"`
	IntervalSeconds int
	gorm.Model
}
