package dbmodels

import (
	"go-web-server-2-practice/internal/gormutil"
)

type User struct {
	gormutil.ModelFields
	Name           string `gorm:"column:name;not null" json:"name"`
	Email          string `gorm:"column:email;unique;uniqueIndex" json:"email"`
	PasswordHash   string `gorm:"column:password_hash;not null" json:"-"`
}