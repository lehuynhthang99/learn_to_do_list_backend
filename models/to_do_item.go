package models

type ToDoListItem struct {
	ID     uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	ListID uint    `gorm:"not null" json:"list_id"`       // Foreign key to ToDoList
	Info   *string `gorm:"type:varchar(100)" json:"info"` // Nullable field
	Status uint8   `gorm:"not null" json:"status"`        // tinyint unsigned → uint8

	// Optional: GORM relationship
	List ToDoList `gorm:"foreignKey:ListID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"list,omitempty"`
}
