package databaseRepository

import (
	"neoway_test/internal/domain/customer/entity"
	"neoway_test/internal/domain/customer/repository"

	"gorm.io/gorm"
)

type CustomerRepositoryPostgres struct {
	Db *gorm.DB
}

func NewPostgresCustomerRepository(db *gorm.DB) (repository.CustomerRepository, error) {
	return &CustomerRepositoryPostgres{Db: db}, nil
}

func (c *CustomerRepositoryPostgres) Create(customer *entity.Customer) error {
	tx := c.Db.Create(customer)
	return tx.Error
}

func (c *CustomerRepositoryPostgres) CreateBulk(customers []*entity.Customer) error {
	tx := c.Db.CreateInBatches(customers, 1000) // Insert in batches of 1000
	return tx.Error
}

func (c *CustomerRepositoryPostgres) Get(page int) ([]*entity.Customer, error) {
	const pageSize = 100
	offset := (page - 1) * pageSize

	var customers []*entity.Customer
	tx := c.Db.Limit(pageSize).Offset(offset).Find(&customers)
	return customers, tx.Error
}

func (c *CustomerRepositoryPostgres) GetById(id string) (*entity.Customer, error) {
	var customer entity.Customer
	tx := c.Db.First(&customer, "id = ?", id)
	return &customer, tx.Error
}

func (c *CustomerRepositoryPostgres) GetByCpf(cpf string) (*entity.Customer, error) {
	var customer entity.Customer
	tx := c.Db.First(&customer, "cpf = ?", cpf)
	return &customer, tx.Error
}

func (c *CustomerRepositoryPostgres) Delete(customer *entity.Customer) error {
	tx := c.Db.Delete(customer)
	return tx.Error
}
