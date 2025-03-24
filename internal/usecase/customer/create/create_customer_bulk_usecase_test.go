package usecase

import (
	"bytes"
	"errors"
	"neoway_test/internal/domain/customer/service"
	databaseRepository "neoway_test/internal/infrastructure/database/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateCustomerBulkUseCase_Success(t *testing.T) {
	fileContent := `CPF                PRIVATE     INCOMPLETO  DATA DA ÚLTIMA COMPRA TICKET MÉDIO          TICKET DA ÚLTIMA COMPRA LOJA MAIS FREQUÊNTE LOJA DA ÚLTIMA COMPRA
026.987.379-13     0           0           2011-01-20            159,31                159,31                  79.379.491/0001-83  79.379.491/0001-83
041.091.641-25     0           1           NULL                  NULL                  NULL                    NULL                NULL`
	reader := bytes.NewReader([]byte(fileContent))

	mockRepo := new(databaseRepository.CustomerRepositoryMock)
	parseService := service.NewParseTxtFileService()
	createCustomerBulkUseCase := NewCreateCustomersBulkUseCase(mockRepo, parseService)

	mockRepo.On("CreateBulk", mock.AnythingOfType("[]*entity.Customer")).Return(nil)

	result, err := createCustomerBulkUseCase.Execute(reader)

	assert.Nil(t, err)
	assert.Equal(t, "Bulk Insert Successful", result)
	mockRepo.AssertExpectations(t)
}

func TestCreateCustomerBulkUseCase_ParsingError(t *testing.T) {
	invalidFileContent := `CPF               PRIVATE        INCOMPLETO     DATA_ULTIMA_COMPRA
922.488.109-20   0              0              2011-01-27`
	reader := bytes.NewReader([]byte(invalidFileContent))

	mockRepo := new(databaseRepository.CustomerRepositoryMock)
	parseService := service.NewParseTxtFileService()
	createCustomerBulkUseCase := NewCreateCustomersBulkUseCase(mockRepo, parseService)

	result, err := createCustomerBulkUseCase.Execute(reader)

	assert.Empty(t, result)
	assert.EqualError(t, err, "invalid file format: line too short")
	mockRepo.AssertNotCalled(t, "CreateBulk")
}

func TestCreateCustomerBulkUseCase_RepositoryError(t *testing.T) {
	fileContent := `CPF                PRIVATE     INCOMPLETO  DATA DA ÚLTIMA COMPRA TICKET MÉDIO          TICKET DA ÚLTIMA COMPRA LOJA MAIS FREQUÊNTE LOJA DA ÚLTIMA COMPRA
	026.987.379-13     0           0           2011-01-20            159.31                159.31                  79.379.491/0001-83  79.379.491/0001-83`
	reader := bytes.NewReader([]byte(fileContent))

	mockRepo := new(databaseRepository.CustomerRepositoryMock)
	parseService := service.NewParseTxtFileService()
	createCustomerBulkUseCase := NewCreateCustomersBulkUseCase(mockRepo, parseService)

	mockRepo.On("CreateBulk", mock.AnythingOfType("[]*entity.Customer")).Return(errors.New("database error"))

	result, err := createCustomerBulkUseCase.Execute(reader)

	assert.Empty(t, result)
	assert.EqualError(t, err, "internal server error")
	mockRepo.AssertExpectations(t)
}
