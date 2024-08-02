package customer

import (
	"github.com/ooo-team/yafds/internal/repository"
	def "github.com/ooo-team/yafds/internal/service"
)

var _ def.CustomerService = (*service)(nil)

type service struct {
	repo repository.CustomerRepository
}

func NewService(
	repo repository.CustomerRepository,
) *service {
	return &service{
		repo: repo,
	}
}
