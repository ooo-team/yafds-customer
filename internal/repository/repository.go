package repository

import (
	"context"

	model "github.com/ooo-team/yafds/internal/model/customer"
)

type CustomerRepository interface {
	Create(ctx context.Context, customerID uint64, info *model.CustomerInfo) error
	Get(ctx context.Context, customerID uint64) (*model.Customer, error)
}
