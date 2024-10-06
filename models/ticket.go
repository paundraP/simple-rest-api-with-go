package models

type Ticket struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Name   string `json:"name"`
	NIK    int64  `json:"nik"`
	Amount int64  `json:"amount"`
	Status bool   `json:"status"`
}
