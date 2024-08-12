package service

import (
	"context"

	model "github.com/ooo-team/yafds-customer/internal/model/customer"
)

type CustomerService interface {
	Create(ctx context.Context, info *model.CustomerInfo) (uint32, error)
	Get(ctx context.Context, uuid uint32, need_metainfo bool) (*model.Customer, error)
}
