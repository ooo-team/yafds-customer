package customer

import (
	"context"
	"log"

	model "github.com/ooo-team/yafds/internal/model/customer"
)

func (s *service) Get(ctx context.Context, uuid uint32, need_metainfo bool) (*model.Customer, error) {
	customer_info, err := s.repo.Get(ctx, uuid)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return customer_info, nil

}
