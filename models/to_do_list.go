package models

import "time"

type ToDoList struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"`       // foreign key to Users
	Name      *string   `gorm:"type:varchar(100)" json:"name"` // nullable string
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`

	// Optional: add relation if using GORM association
	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user,omitempty"`
}
