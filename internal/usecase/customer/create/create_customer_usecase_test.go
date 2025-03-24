package usecase

import (
	"errors"
	"neoway_test/internal/domain/customer/dto"
	"neoway_test/internal/domain/customer/entity"
	"neoway_test/internal/domain/customer/service"
	databaseRepository "neoway_test/internal/infrastructure/database/repository"
	internalerrors "neoway_test/internal/internal-errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateCustomerUseCase_Success(t *testing.T) {
	mockRepo := new(databaseRepository.CustomerRepositoryMock)
	parseService := service.NewParseService()
	createCustomerUseCase := NewCreateCustomerUseCase(mockRepo, parseService)

	dataUltimaCompra := time.Date(2011, 10, 4, 0, 0, 0, 0, time.UTC)

	input := dto.InputCreateCustomerDto{
		Cpf:                "152.298.818-10",
		Private:            "0",
		Incompleto:         "1",
		DataUltimaCompra:   "2011-10-04",
		TicketMedio:        100.5,
		TicketUltimaCompra: 200.75,
		LojaMaisFrequente:  "79.379.491/0008-50",
		LojaUltimaCompra:   "79.379.491/0008-50",
	}

	customer := &entity.Customer{
		Cpf:                         input.Cpf,
		CpfValido:                   true,
		Private:                     input.Private,
		Incompleto:                  input.Incompleto,
		DataUltimaCompra:            &dataUltimaCompra,
		TicketMedio:                 input.TicketMedio,
		TicketUltimaCompra:          input.TicketUltimaCompra,
		LojaMaisFrequente:           input.LojaMaisFrequente,
		CnpjLojaMaisFrequenteValido: true,
		LojaUltimaCompra:            input.LojaUltimaCompra,
		CnpjLojaUltimaCompraValido:  true,
	}

	mockRepo.On("Create", mock.AnythingOfType("*entity.Customer")).Return(nil)

	output, err := createCustomerUseCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, customer.Cpf, output.Cpf)
	assert.Equal(t, customer.Private, output.Private)
	assert.Equal(t, customer.Incompleto, output.Incompleto)
	assert.WithinDuration(t, customer.DataUltimaCompra.UTC(), output.DataUltimaCompra.UTC(), time.Second)
	assert.Equal(t, customer.TicketMedio, output.TicketMedio)
	assert.Equal(t, customer.TicketUltimaCompra, output.TicketUltimaCompra)
	assert.Equal(t, customer.LojaMaisFrequente, output.LojaMaisFrequente)
	assert.Equal(t, customer.LojaUltimaCompra, output.LojaUltimaCompra)
	assert.Equal(t, customer.CpfValido, output.CpfValido)
	assert.Equal(t, customer.CnpjLojaMaisFrequenteValido, output.CnpjLojaMaisFrequenteValido)
	assert.Equal(t, customer.CnpjLojaUltimaCompraValido, output.CnpjLojaUltimaCompraValido)

	mockRepo.AssertExpectations(t)
}

func TestCreateCustomerUseCase_ErrorCreatingCustomer(t *testing.T) {
	mockRepo := new(databaseRepository.CustomerRepositoryMock)
	parseService := service.NewParseService()
	createCustomerUseCase := NewCreateCustomerUseCase(mockRepo, parseService)

	input := dto.InputCreateCustomerDto{
		Cpf: "152.298.818-10",
	}

	mockRepo.On("Create", mock.AnythingOfType("*entity.Customer")).Return(errors.New("database error"))

	output, err := createCustomerUseCase.Execute(input)

	assert.Empty(t, output)
	assert.Equal(t, internalerrors.ErrInternal, err)
	mockRepo.AssertExpectations(t)
}

func TestCreateCustomerUseCase_EmptyCustomerData(t *testing.T) {
	mockRepo := new(databaseRepository.CustomerRepositoryMock)
	parseService := service.NewParseService()
	createCustomerUseCase := NewCreateCustomerUseCase(mockRepo, parseService)

	input := dto.InputCreateCustomerDto{
		Cpf:                "",
		LojaMaisFrequente:  "",
		LojaUltimaCompra:   "",
		DataUltimaCompra:   "NULL",
		TicketMedio:        0.0,
		TicketUltimaCompra: 0.0,
	}

	expectedCustomer := &entity.Customer{
		Cpf:                         "NULL",
		CpfValido:                   false,
		LojaMaisFrequente:           "NULL",
		LojaUltimaCompra:            "NULL",
		DataUltimaCompra:            nil,
		TicketMedio:                 0.0,
		TicketUltimaCompra:          0.0,
		CnpjLojaMaisFrequenteValido: false,
		CnpjLojaUltimaCompraValido:  false,
	}

	mockRepo.On("Create", mock.AnythingOfType("*entity.Customer")).Return(nil)

	output, err := createCustomerUseCase.Execute(input)

	assert.NotEmpty(t, output)
	assert.Nil(t, err)
	assert.Equal(t, expectedCustomer.Cpf, output.Cpf)
	assert.False(t, output.CpfValido)
	assert.Equal(t, expectedCustomer.LojaMaisFrequente, output.LojaMaisFrequente)
	assert.Equal(t, expectedCustomer.LojaUltimaCompra, output.LojaUltimaCompra)
	assert.Nil(t, output.DataUltimaCompra)
	assert.Equal(t, expectedCustomer.TicketMedio, output.TicketMedio)
	assert.Equal(t, expectedCustomer.TicketUltimaCompra, output.TicketUltimaCompra)
	assert.False(t, output.CnpjLojaMaisFrequenteValido)
	assert.False(t, output.CnpjLojaUltimaCompraValido)
	assert.Equal(t, expectedCustomer.CpfValido, output.CpfValido)
	assert.Equal(t, expectedCustomer.CnpjLojaMaisFrequenteValido, output.CnpjLojaMaisFrequenteValido)
	assert.Equal(t, expectedCustomer.CnpjLojaUltimaCompraValido, output.CnpjLojaUltimaCompraValido)
	assert.Equal(t, expectedCustomer.DataUltimaCompra, output.DataUltimaCompra)

	mockRepo.AssertExpectations(t)
}
