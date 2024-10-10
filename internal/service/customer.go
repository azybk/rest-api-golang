package service

import (
	"context"
	"database/sql"
	"rest-api-golang/domain"
	"rest-api-golang/dto"
	"time"

	"github.com/google/uuid"
)

type customerService struct {
	customerRepository domain.CustomerRepository
}

func NewCustomer(customerRepository domain.CustomerRepository) domain.CustomerService {
	return &customerService{
		customerRepository: customerRepository,
	}
}

func (cs customerService) Index(ctx context.Context) ([]dto.CustomerData, error) {
	customers, err := cs.customerRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var customerData []dto.CustomerData
	for _, v := range customers {
		customerData = append(customerData, dto.CustomerData{
			ID: v.ID,
			Code: v.Code,
			Name: v.Name,
		})
	}

	return customerData, nil
}

func (cs customerService) Create(ctx context.Context, req dto.CreateCustomerRequest) error {
	customer := domain.Customer{
		ID: uuid.NewString(),
		Code: req.Code,
		Name: req.Name,
		CreatedAt: sql.NullTime{Valid: true, Time: time.Now()},
	}

	return cs.customerRepository.Save(ctx, &customer)
}