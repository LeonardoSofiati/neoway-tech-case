package usecase

import (
	"errors"
	"neoway_test/internal/domain/customer/dto"
	"neoway_test/internal/domain/customer/entity"
	shared "neoway_test/internal/domain/shared/entity"
	databaseRepository "neoway_test/internal/infrastructure/database/repository"
	internalerrors "neoway_test/internal/internal-errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetCustomerByIdUseCase_Success(t *testing.T) {
	mockRepo := new(databaseRepository.CustomerRepositoryMock)
	getCustomerByIdUseCase := NewGetCustomerByIdUseCase(mockRepo)

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

	input := dto.InputGetCustomerByIdDto{
		ID: customer.ID,
	}

	mockRepo.On("GetById", input.ID).Return(customer, nil)

	output, err := getCustomerByIdUseCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, customer.ID, output.ID)
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

func TestGetCustomerByIdUseCase_CustomerNotFound(t *testing.T) {
	mockRepo := new(databaseRepository.CustomerRepositoryMock)
	getCustomerByIdUseCase := NewGetCustomerByIdUseCase(mockRepo)

	input := dto.InputGetCustomerByIdDto{
		ID: "customer123",
	}

	mockRepo.On("GetById", input.ID).Return(nil, errors.New("record not found"))

	output, err := getCustomerByIdUseCase.Execute(input)

	assert.Nil(t, output)
	assert.Equal(t, internalerrors.ProcessErrorToReturn(errors.New("record not found")), err)
	mockRepo.AssertExpectations(t)
}

func TestGetCustomerByIdUseCase_InternalError(t *testing.T) {
	mockRepo := new(databaseRepository.CustomerRepositoryMock)
	getCustomerByIdUseCase := NewGetCustomerByIdUseCase(mockRepo)

	input := dto.InputGetCustomerByIdDto{
		ID: "customer123",
	}

	mockRepo.On("GetById", input.ID).Return(nil, internalerrors.ErrInternal)

	output, err := getCustomerByIdUseCase.Execute(input)

	assert.Nil(t, output)
	assert.Equal(t, internalerrors.ErrInternal, err)
	mockRepo.AssertExpectations(t)
}
