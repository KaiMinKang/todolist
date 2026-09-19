package main

import (
	"time"
)

type UserModel struct {
	ID           int64     `gorm:"primaryKey;aotoIncrement" json:"id"`
	Username     string    `grom:"size(32);uniqueIndex;not null" json:"username"`
	PasswordHash string    `grom:"size(60);not null" json:"-"`
	CreatedAt    time.Time `grom:"precision:3" json:"created_at"`
}
type TodoModel struct {
	UserId    int64     `grom:"primaryKey;autoIncrement" json:"user_id"`
	Title     string    `grom:"not null;index" json:"title"`
	Done      bool      `grom:"not null;default:flase" json:"done"`
	CreatedAt time.Time `gorm:"precision:3" json:"created_at"`
	UpdatedAt time.Time `gorm:"precision:3" json:"updated_at"`
}

func (UserModel) TableName() string { return "users" }
func (TodoModel) TableName() string { return "todos" }
