package usecase

import (
	"io"
	"neoway_test/internal/domain/customer/entity"
	"neoway_test/internal/domain/customer/repository"
	"neoway_test/internal/domain/customer/service"
	internalerrors "neoway_test/internal/internal-errors"
)

type CreateCustomerBulkUseCase struct {
	repo                repository.CustomerRepository
	parseTxtFileService *service.ParseTxtFileService
}

func NewCreateCustomersBulkUseCase(repo repository.CustomerRepository, parseService *service.ParseTxtFileService) *CreateCustomerBulkUseCase {
	return &CreateCustomerBulkUseCase{
		repo:                repo,
		parseTxtFileService: parseService,
	}
}

func (uc *CreateCustomerBulkUseCase) Execute(file io.Reader) (string, error) {
	customersDTO, err := uc.parseTxtFileService.ExecuteParseTxtFileService(file)

	if err != nil {
		return "", err
	}

	var customers []*entity.Customer
	for _, newCustomer := range customersDTO {
		customer, err := entity.NewCustomer(
			newCustomer.Cpf,
			newCustomer.Private,
			newCustomer.Incompleto,
			newCustomer.DataUltimaCompra,
			newCustomer.TicketMedio,
			newCustomer.TicketUltimaCompra,
			newCustomer.LojaMaisFrequente,
			newCustomer.LojaUltimaCompra,
		)
		if err != nil {
			return "", err
		}
		customers = append(customers, customer)
	}

	// 3. Salva no repositório
	err = uc.repo.CreateBulk(customers)
	if err != nil {
		return "", internalerrors.ErrInternal
	}

	return "Bulk Insert Successful", nil
}
