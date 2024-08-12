package converter

import (
	"time"

	model "github.com/ooo-team/yafds-customer/internal/model/customer"
	repoModel "github.com/ooo-team/yafds-customer/internal/repository/customer/model"
)

func ToCustomerFromRepo(customer *repoModel.Customer) *model.Customer {

	var updatedAt time.Time
	if customer.UpdatedAt.Valid {
		updatedAt = customer.UpdatedAt.Time
	}

	return &model.Customer{
		ID:        customer.ID,
		Info:      ToCustomerInfoFromRepo(customer.Info),
		CreatedAt: &customer.CreatedAt,
		UpdatedAt: &updatedAt,
	}
}

func ToCustomerInfoFromRepo(info repoModel.CustomerInfo) model.CustomerInfo {
	return model.CustomerInfo{
		Phone:   info.Phone,
		Email:   info.Email,
		Address: info.Address,
	}
}

func ToCustomerInfoFromService(info *model.CustomerInfo) repoModel.CustomerInfo {
	return repoModel.CustomerInfo{
		Phone:   info.Phone,
		Email:   info.Email,
		Address: info.Address,
	}
}
