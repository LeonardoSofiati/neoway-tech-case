package databaseRepository

import (
	"neoway_test/internal/domain/customer/entity"

	"github.com/stretchr/testify/mock"
)

type CustomerRepositoryMock struct {
	mock.Mock
}

func (r *CustomerRepositoryMock) Create(customer *entity.Customer) error {
	args := r.Called(customer)
	return args.Error(0)
}

func (r *CustomerRepositoryMock) CreateBulk(customers []*entity.Customer) error {
	args := r.Called(customers)
	return args.Error(0)
}

func (r *CustomerRepositoryMock) Get(page int) ([]*entity.Customer, error) {
	args := r.Called(page)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*entity.Customer), nil
}

func (r *CustomerRepositoryMock) GetById(id string) (*entity.Customer, error) {
	args := r.Called(id)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Customer), nil
}

func (r *CustomerRepositoryMock) GetByCpf(cpf string) (*entity.Customer, error) {
	args := r.Called(cpf)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Customer), nil
}

func (r *CustomerRepositoryMock) Delete(customer *entity.Customer) error {
	args := r.Called(customer)
	return args.Error(0)
}
