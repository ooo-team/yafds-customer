package customer

import (
	"context"
	"log"

	"github.com/google/uuid"

	model "github.com/ooo-team/yafds/internal/model/customer"
)

func (s *service) Create(ctx context.Context, info *model.CustomerInfo) (string, error) {
	userUUID, err := uuid.NewUUID()
	if err != nil {
		log.Printf("Failed to generate uuid: %v\n", err)
		return "Failed to generate uuid", err
	}
	err = s.repo.Create(ctx, uint64(userUUID.ID()), info)
	if err != nil {
		log.Printf("Failed to create user: %v\n", err)
		return "Failed to create user", err
	}

	return userUUID.String(), nil
}
