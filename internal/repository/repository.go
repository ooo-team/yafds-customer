package repository

import (
	"context"

	model "github.com/ooo-team/yafds-customer/internal/model/customer"
)

type CustomerRepository interface {
	Create(ctx context.Context, customerID uint32, info *model.CustomerInfo) error
	Get(ctx context.Context, customerID uint32) (*model.Customer, error)
}
