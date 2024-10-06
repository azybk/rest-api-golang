package repository

import (
	"context"
	"database/sql"
	"rest-api-golang/domain"

	"github.com/doug-martin/goqu/v9"
)

type customerRepository struct {
	db *goqu.Database
}

func NewCustomer(con *sql.DB) domain.CustomerRepository {
	return &customerRepository{
		db: goqu.New("default", con),
	}
}

func (cr customerRepository) FindAll(ctx context.Context) ([]domain.Customer, error) {
	panic("implement me")
}

func (cr customerRepository) FindById(ctx context.Context, id string) (domain.Customer, error) {
	panic("implement me")
}

func (cr customerRepository) Save(ctx context.Context, c *domain.Customer) error {
	panic("implement me")
}

func (cr customerRepository) Update(ctx context.Context, c *domain.Customer) error {
	panic("implement me")
}

func (cr customerRepository) Delete(ctx context.Context, id string) error {
	panic("implement me")
}
