package app

import (
	"github.com/ooo-team/yafds/internal/repository"
	customerRepository "github.com/ooo-team/yafds/internal/repository/customer"
	"github.com/ooo-team/yafds/internal/service"
	customerService "github.com/ooo-team/yafds/internal/service/customer"
)

type serviceProvider struct {
	customerService service.CustomerService
	customerRepo    repository.CustomerRepository
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) CustomerRepo() repository.CustomerRepository {
	if s.customerRepo == nil {
		s.customerRepo = customerRepository.NewRepository()
	}
	return s.customerRepo
}

func (s *serviceProvider) CustomerService() service.CustomerService {
	if s.customerService == nil {
		s.customerService = customerService.NewService(s.CustomerRepo())
	}
	return s.customerService
}
