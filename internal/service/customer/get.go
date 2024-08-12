package customer

import (
	"context"
	"log"

	model "github.com/ooo-team/yafds-customer/internal/model/customer"
)

func (s *service) Get(ctx context.Context, uuid uint32, needMetainfo bool) (*model.Customer, error) {
	customerInfo, err := s.repo.Get(ctx, uuid)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return customerInfo, nil

}
