package user

import "time"

type ID uint

type User struct {
	ID        ID         `gorm:"primary_key" json:"id"`
	Name      string     `json:"name"`
	Age       string     `json:"age"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
