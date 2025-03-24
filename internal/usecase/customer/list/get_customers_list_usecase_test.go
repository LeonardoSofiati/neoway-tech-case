package usecase

import (
	"neoway_test/internal/domain/customer/dto"
	"neoway_test/internal/domain/customer/entity"
	shared "neoway_test/internal/domain/shared/entity"
	databaseRepository "neoway_test/internal/infrastructure/database/repository"
	internalerrors "neoway_test/internal/internal-errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetCustomersListUseCase_Success(t *testing.T) {
	mockRepo := new(databaseRepository.CustomerRepositoryMock)
	getCustomersListUseCase := NewGetCustomersListUseCase(mockRepo)

	dataUltimaCompra := time.Date(2011, 10, 5, 0, 0, 0, 0, time.UTC)

	customers := []*entity.Customer{
		{
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
		},
	}

	input := dto.InputGetCustomersListDto{Page: 1}

	mockRepo.On("Get", input.Page).Return(customers, nil)

	output, err := getCustomersListUseCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Len(t, output, 1)
	assert.Equal(t, customers[0].Cpf, output[0].Cpf)
	assert.Equal(t, customers[0].Private, output[0].Private)
	assert.Equal(t, customers[0].Incompleto, output[0].Incompleto)
	assert.WithinDuration(t, customers[0].DataUltimaCompra.UTC(), output[0].DataUltimaCompra.UTC(), time.Second)
	assert.Equal(t, customers[0].TicketMedio, output[0].TicketMedio)
	assert.Equal(t, customers[0].TicketUltimaCompra, output[0].TicketUltimaCompra)
	assert.Equal(t, customers[0].LojaMaisFrequente, output[0].LojaMaisFrequente)
	assert.Equal(t, customers[0].LojaUltimaCompra, output[0].LojaUltimaCompra)
	assert.Equal(t, customers[0].CpfValido, output[0].CpfValido)
	assert.Equal(t, customers[0].CnpjLojaMaisFrequenteValido, output[0].CnpjLojaMaisFrequenteValido)
	assert.Equal(t, customers[0].CnpjLojaUltimaCompraValido, output[0].CnpjLojaUltimaCompraValido)

	mockRepo.AssertExpectations(t)
}

func TestGetCustomersListUseCase_EmptyList(t *testing.T) {
	mockRepo := new(databaseRepository.CustomerRepositoryMock)
	getCustomersListUseCase := NewGetCustomersListUseCase(mockRepo)

	input := dto.InputGetCustomersListDto{Page: 1}

	mockRepo.On("Get", input.Page).Return([]*entity.Customer{}, nil)

	output, err := getCustomersListUseCase.Execute(input)

	assert.Nil(t, err)
	assert.Nil(t, output)
	assert.Len(t, output, 0)

	mockRepo.AssertExpectations(t)
}

func TestGetCustomersListUseCase_InternalError(t *testing.T) {
	mockRepo := new(databaseRepository.CustomerRepositoryMock)
	getCustomersListUseCase := NewGetCustomersListUseCase(mockRepo)

	input := dto.InputGetCustomersListDto{Page: 1}

	mockRepo.On("Get", input.Page).Return(nil, internalerrors.ErrInternal)

	output, err := getCustomersListUseCase.Execute(input)

	assert.Nil(t, output)
	assert.Equal(t, internalerrors.ErrInternal, err)

	mockRepo.AssertExpectations(t)
}
