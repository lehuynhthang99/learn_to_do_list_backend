package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

const (
	NotStarted = iota
	InProgress
	Done
)

type ToDoItem struct {
	Info   *string `json:"info"`
	Status uint8   `json:"status"`
}

type ToDoItems []ToDoItem

// save as JSON to DB
func (items ToDoItems) Value() (driver.Value, error) {
	return json.Marshal(items)
}

// read from JSON in DB
func (t *ToDoItems) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan ToDoItems value: %v", value)
	}

	return json.Unmarshal(bytes, &t)
}
