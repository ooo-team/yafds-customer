package repository

import (
	model "github.com/ooo-team/yafds/internal/model/customer"
)

type CustomerRepository interface {
	Create(userID uint64, info *model.CustomerInfo) error
	Get(userID uint64) (*model.Customer, error)
}
