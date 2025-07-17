package models

import "time"

type ToDoList struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"`       // foreign key to Users
	Name      *string   `gorm:"type:varchar(100)" json:"name"` // nullable string
	Items     ToDoItems `gorm:"type:json" json:"items"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (ToDoList) TableName() string {
	return "to_do_list_tb"
}
