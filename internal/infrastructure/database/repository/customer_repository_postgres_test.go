package databaseRepository_test

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"neoway_test/internal/domain/customer/entity"
	shared "neoway_test/internal/domain/shared/entity"
	databaseRepository "neoway_test/internal/infrastructure/database/repository"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	var err error
	dsn := "host=localhost user=neoway_dev password=password dbname=neoway_dev port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	db.AutoMigrate(&entity.Customer{})

	code := m.Run()
	os.Exit(code)
}

func setupTestDB() {
	db.Exec("DROP TABLE IF EXISTS customers")
	db.AutoMigrate(&entity.Customer{})
}

func TestPostgresCustomerRepository(t *testing.T) {
	repo, _ := databaseRepository.NewPostgresCustomerRepository(db)

	t.Run("Create", func(t *testing.T) {
		setupTestDB()
		dataUltimaCompra := time.Date(2011, 10, 5, 0, 0, 0, 0, time.UTC)

		customer := &entity.Customer{
			BaseEntity:                  shared.NewBaseEntity(),
			Cpf:                         "922.488.109-20",
			CpfValido:                   true,
			Private:                     "1",
			Incompleto:                  "0",
			DataUltimaCompra:            &dataUltimaCompra,
			TicketMedio:                 130.54,
			TicketUltimaCompra:          130.54,
			LojaMaisFrequente:           "79.379.491/0001-83",
			CnpjLojaMaisFrequenteValido: true,
			LojaUltimaCompra:            "79.379.491/0001-83",
			CnpjLojaUltimaCompraValido:  true,
		}
		err := repo.Create(customer)
		assert.Nil(t, err)

		storedCustomer, err := repo.GetById(customer.ID)

		assert.Nil(t, err)
		assert.Equal(t, customer.ID, storedCustomer.ID)
		assert.NotNil(t, storedCustomer)
		assert.Equal(t, customer.Cpf, storedCustomer.Cpf)
		assert.Equal(t, customer.Private, storedCustomer.Private)
		assert.Equal(t, customer.Incompleto, storedCustomer.Incompleto)
		assert.Equal(t, customer.DataUltimaCompra.UTC(), storedCustomer.DataUltimaCompra.UTC())
		assert.Equal(t, customer.TicketMedio, storedCustomer.TicketMedio)
		assert.Equal(t, customer.TicketUltimaCompra, storedCustomer.TicketUltimaCompra)
		assert.Equal(t, customer.LojaMaisFrequente, storedCustomer.LojaMaisFrequente)
		assert.Equal(t, customer.CnpjLojaMaisFrequenteValido, storedCustomer.CnpjLojaMaisFrequenteValido)
		assert.Equal(t, customer.LojaUltimaCompra, storedCustomer.LojaUltimaCompra)
		assert.Equal(t, customer.CnpjLojaUltimaCompraValido, storedCustomer.CnpjLojaUltimaCompraValido)
	})

	t.Run("CreateBulk", func(t *testing.T) {
		setupTestDB()
		dataUltimaCompra := time.Date(2011, 10, 5, 0, 0, 0, 0, time.UTC)

		customers := []*entity.Customer{
			{
				BaseEntity:                  shared.NewBaseEntity(),
				Cpf:                         "891.098.302-78",
				CpfValido:                   true,
				Private:                     "1",
				Incompleto:                  "0",
				DataUltimaCompra:            &dataUltimaCompra,
				TicketMedio:                 130.54,
				TicketUltimaCompra:          130.54,
				LojaMaisFrequente:           "79.379.491/0001-83",
				CnpjLojaMaisFrequenteValido: true,
				LojaUltimaCompra:            "79.379.491/0001-83",
				CnpjLojaUltimaCompraValido:  true,
			},
			{
				BaseEntity:                  shared.NewBaseEntity(),
				Cpf:                         "046.857.249-09",
				CpfValido:                   true,
				Private:                     "0",
				Incompleto:                  "1",
				DataUltimaCompra:            &dataUltimaCompra,
				TicketMedio:                 130.54,
				TicketUltimaCompra:          130.54,
				LojaMaisFrequente:           "79.379.491/0001-83",
				CnpjLojaMaisFrequenteValido: true,
				LojaUltimaCompra:            "79.379.491/0001-83",
				CnpjLojaUltimaCompraValido:  true,
			}}
		err := repo.CreateBulk(customers)
		assert.Nil(t, err)

		storedCustomers, err := repo.Get(1)
		assert.Nil(t, err)
		assert.Equal(t, 2, len(storedCustomers))

		storedCustomer, err := repo.GetById(customers[0].ID)

		assert.Nil(t, err)
		assert.Equal(t, customers[0].ID, storedCustomer.ID)
		assert.NotNil(t, storedCustomer)
		assert.Equal(t, customers[0].Cpf, storedCustomer.Cpf)
		assert.Equal(t, customers[0].Private, storedCustomer.Private)
		assert.Equal(t, customers[0].Incompleto, storedCustomer.Incompleto)
		assert.Equal(t, customers[0].DataUltimaCompra.UTC(), storedCustomer.DataUltimaCompra.UTC())
		assert.Equal(t, customers[0].TicketMedio, storedCustomer.TicketMedio)
		assert.Equal(t, customers[0].TicketUltimaCompra, storedCustomer.TicketUltimaCompra)
		assert.Equal(t, customers[0].LojaMaisFrequente, storedCustomer.LojaMaisFrequente)
		assert.Equal(t, customers[0].CnpjLojaMaisFrequenteValido, storedCustomer.CnpjLojaMaisFrequenteValido)
		assert.Equal(t, customers[0].LojaUltimaCompra, storedCustomer.LojaUltimaCompra)
		assert.Equal(t, customers[0].CnpjLojaUltimaCompraValido, storedCustomer.CnpjLojaUltimaCompraValido)

		storedCustomer2, err := repo.GetById(customers[1].ID)

		assert.Nil(t, err)
		assert.Equal(t, customers[1].ID, storedCustomer2.ID)
		assert.NotNil(t, storedCustomer2)
		assert.Equal(t, customers[1].Cpf, storedCustomer2.Cpf)
		assert.Equal(t, customers[1].Private, storedCustomer2.Private)
		assert.Equal(t, customers[1].Incompleto, storedCustomer2.Incompleto)
		assert.Equal(t, customers[1].DataUltimaCompra.UTC(), storedCustomer2.DataUltimaCompra.UTC())
		assert.Equal(t, customers[1].TicketMedio, storedCustomer2.TicketMedio)
		assert.Equal(t, customers[1].TicketUltimaCompra, storedCustomer2.TicketUltimaCompra)
		assert.Equal(t, customers[1].LojaMaisFrequente, storedCustomer2.LojaMaisFrequente)
		assert.Equal(t, customers[1].CnpjLojaMaisFrequenteValido, storedCustomer2.CnpjLojaMaisFrequenteValido)
		assert.Equal(t, customers[1].LojaUltimaCompra, storedCustomer2.LojaUltimaCompra)
		assert.Equal(t, customers[1].CnpjLojaUltimaCompraValido, storedCustomer2.CnpjLojaUltimaCompraValido)
	})

	t.Run("GetByCpf", func(t *testing.T) {
		setupTestDB()

		dataUltimaCompra := time.Date(2011, 10, 5, 0, 0, 0, 0, time.UTC)

		customer := &entity.Customer{
			BaseEntity:                  shared.NewBaseEntity(),
			Cpf:                         "922.488.109-20",
			CpfValido:                   true,
			Private:                     "1",
			Incompleto:                  "0",
			DataUltimaCompra:            &dataUltimaCompra,
			TicketMedio:                 130.54,
			TicketUltimaCompra:          130.54,
			LojaMaisFrequente:           "79.379.491/0001-83",
			CnpjLojaMaisFrequenteValido: true,
			LojaUltimaCompra:            "79.379.491/0001-83",
			CnpjLojaUltimaCompraValido:  true,
		}
		repo.Create(customer)

		storedCustomer, err := repo.GetByCpf(customer.Cpf)
		assert.Nil(t, err)
		assert.Equal(t, customer.ID, storedCustomer.ID)
	})

	t.Run("GetById", func(t *testing.T) {
		setupTestDB()

		dataUltimaCompra := time.Date(2011, 10, 5, 0, 0, 0, 0, time.UTC)

		customer := &entity.Customer{
			BaseEntity:                  shared.NewBaseEntity(),
			Cpf:                         "922.488.109-20",
			CpfValido:                   true,
			Private:                     "1",
			Incompleto:                  "0",
			DataUltimaCompra:            &dataUltimaCompra,
			TicketMedio:                 130.54,
			TicketUltimaCompra:          130.54,
			LojaMaisFrequente:           "79.379.491/0001-83",
			CnpjLojaMaisFrequenteValido: true,
			LojaUltimaCompra:            "79.379.491/0001-83",
			CnpjLojaUltimaCompraValido:  true,
		}
		repo.Create(customer)

		storedCustomer, err := repo.GetById(customer.ID)
		assert.Nil(t, err)
		assert.Equal(t, customer.ID, storedCustomer.ID)
	})

	t.Run("Delete", func(t *testing.T) {
		setupTestDB()

		dataUltimaCompra := time.Date(2011, 10, 5, 0, 0, 0, 0, time.UTC)

		customer := &entity.Customer{
			BaseEntity:                  shared.NewBaseEntity(),
			Cpf:                         "922.488.109-20",
			CpfValido:                   true,
			Private:                     "1",
			Incompleto:                  "0",
			DataUltimaCompra:            &dataUltimaCompra,
			TicketMedio:                 130.54,
			TicketUltimaCompra:          130.54,
			LojaMaisFrequente:           "79.379.491/0001-83",
			CnpjLojaMaisFrequenteValido: true,
			LojaUltimaCompra:            "79.379.491/0001-83",
			CnpjLojaUltimaCompraValido:  true,
		}
		repo.Create(customer)

		err := repo.Delete(customer)
		assert.Nil(t, err)

		storedCustomer, err := repo.GetById(customer.ID)

		assert.Error(t, err)
		assert.NotNil(t, storedCustomer)
		assert.Equal(t, gorm.ErrRecordNotFound, err)
	})
}
