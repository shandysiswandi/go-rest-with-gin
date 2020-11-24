package models

import (
	"github.com/shandysiswandi/go-rest-with-gin/utils"
	"gorm.io/gorm"
)

// User is ..
type User struct {
	Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// TableName is
func (u *User) TableName() string {
	return "users"
}

// BeforeCreate is
func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = utils.GenUUID()
	return nil
}
