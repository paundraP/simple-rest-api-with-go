package models

import (
	"time"
)

type User struct {
	ID            uint      `json:"id" gorm:"primary_key"`
	Email         string    `json:"email" gorm:"unique"`
	Password      string    `json:"password"`
	Role          string    `json:"role"`
	OrderedTicket bool      `json:"orderedticket"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
