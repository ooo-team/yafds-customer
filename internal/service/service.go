package service

import (
	"context"

	model "github.com/ooo-team/yafds/internal/model/customer"
)

type CustomerService interface {
	Create(ctx context.Context, info *model.CustomerInfo) (string, error)
	Get(ctx context.Context, uuid string) (*model.Customer, error)
}
