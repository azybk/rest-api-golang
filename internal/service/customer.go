package service

import (
	"context"
	"database/sql"
	"errors"
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
			ID:   v.ID,
			Code: v.Code,
			Name: v.Name,
		})
	}

	return customerData, nil
}

func (cs customerService) Create(ctx context.Context, req dto.CreateCustomerRequest) error {
	customer := domain.Customer{
		ID:        uuid.NewString(),
		Code:      req.Code,
		Name:      req.Name,
		CreatedAt: sql.NullTime{Valid: true, Time: time.Now()},
	}

	return cs.customerRepository.Save(ctx, &customer)
}

func (cs customerService) Update(ctx context.Context, req dto.UpdateCustomerRequest) error {
	persisted, err := cs.customerRepository.FindById(ctx, req.ID)
	if err != nil {
		return err
	}

	if persisted.ID == "" {
		return errors.New("data customer tidak ditemukan")
	}

	persisted.Code = req.Code
	persisted.Name = req.Name
	persisted.UpdatedAt = sql.NullTime{Valid: true, Time: time.Now()}

	return cs.customerRepository.Update(ctx, &persisted)
}

func (cs customerService) Delete(ctx context.Context, id string) error {
	exist, err := cs.customerRepository.FindById(ctx, id)
	if err != nil {
		return err
	}

	if exist.ID == "" {
		return errors.New("data customer tidak ditemukan")
	}

	return cs.customerRepository.Delete(ctx, id)
}

func (cs customerService) Show(ctx context.Context, id string) (dto.CustomerData, error) {
	persisted, err := cs.customerRepository.FindById(ctx, id)
	if err != nil {
		return dto.CustomerData{}, err
	}

	if persisted.ID == "" {
		return dto.CustomerData{}, errors.New("data customer tidak ditemukan")
	}

	return dto.CustomerData{
		ID:   persisted.ID,
		Code: persisted.Code,
		Name: persisted.Name,
	}, nil
}
