package repository

import (
	"neoway_test/internal/domain/customer/entity"
	shared "neoway_test/internal/domain/shared/repository"
)

type CustomerRepository interface {
	shared.RepositoryInterface[entity.Customer]
	GetByCpf(cpf string) (*entity.Customer, error)
	CreateBulk(customers []*entity.Customer) error
}
