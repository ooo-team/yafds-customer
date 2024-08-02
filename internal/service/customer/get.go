package customer

import (
	"context"
	"log"

	model "github.com/ooo-team/yafds/internal/model/customer"
)

func (s *service) Get(ctx context.Context, uuid uint32, need_metainfo bool) (*model.Customer, error) {
	customer_info, err := s.repo.Get(ctx, uuid, need_metainfo)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &model.Customer{
		ID:        uuid,
		Info:      *customer_info,
		CreatedAt: nil,
		UpdatedAt: nil,
	}, nil

}
