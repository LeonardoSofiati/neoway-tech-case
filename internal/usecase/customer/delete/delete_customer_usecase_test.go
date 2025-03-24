package usecase

import (
	"errors"
	"neoway_test/internal/domain/customer/dto"
	"neoway_test/internal/domain/customer/entity"
	shared "neoway_test/internal/domain/shared/entity"
	databaseRepository "neoway_test/internal/infrastructure/database/repository"
	internalerrors "neoway_test/internal/internal-errors"
	"time"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeleteCustomerUseCase_Success(t *testing.T) {
	mockRepo := new(databaseRepository.CustomerRepositoryMock)
	deleteCustomerUseCase := NewDeleteCustomerUseCase(mockRepo)

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

	input := dto.InputDeleteCustomerDto{
		ID: customer.ID,
	}

	mockRepo.On("GetById", input.ID).Return(customer, nil)
	mockRepo.On("Delete", mock.AnythingOfType("*entity.Customer")).Return(nil)

	err := deleteCustomerUseCase.Execute(input)

	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteCustomerUseCase_CustomerNotFound(t *testing.T) {
	mockRepo := new(databaseRepository.CustomerRepositoryMock)
	deleteCustomerUseCase := NewDeleteCustomerUseCase(mockRepo)

	input := dto.InputDeleteCustomerDto{
		ID: "customer123",
	}

	mockRepo.On("GetById", input.ID).Return(nil, errors.New("record not found"))

	err := deleteCustomerUseCase.Execute(input)

	assert.Equal(t, errors.New("record not found"), err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteCustomerUseCase_ErrorDeletingCustomer(t *testing.T) {
	mockRepo := new(databaseRepository.CustomerRepositoryMock)
	deleteCustomerUseCase := NewDeleteCustomerUseCase(mockRepo)
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

	input := dto.InputDeleteCustomerDto{
		ID: customer.ID,
	}

	mockRepo.On("GetById", input.ID).Return(customer, nil)

	mockRepo.On("Delete", mock.AnythingOfType("*entity.Customer")).Return(internalerrors.ErrInternal)

	err := deleteCustomerUseCase.Execute(input)

	assert.Equal(t, internalerrors.ErrInternal, err)
	mockRepo.AssertExpectations(t)
}
