package gormutil

import (
	"gorm.io/gorm"
	"time"
)

type Model interface {
	mustEmbeModelFields()
}

type ModelFields struct {
	ID        uint `gorm:"column:id;primarykey"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (ModelFields) mustEmbeModelFields(){}