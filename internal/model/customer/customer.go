package model

import (
	"time"
)

type Customer struct {
	ID        uint32
	Info      CustomerInfo
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type CustomerInfo struct {
	Phone   string
	Email   string
	Address string
}
