package usecase

import (
	"neoway_test/internal/domain/customer/dto"
	"neoway_test/internal/domain/customer/entity"
	"neoway_test/internal/domain/customer/repository"
	"neoway_test/internal/domain/customer/service"
	internalerrors "neoway_test/internal/internal-errors"
)

type CreateCustomerUseCase struct {
	repo         repository.CustomerRepository
	parseService *service.ParseService
}

func NewCreateCustomerUseCase(repo repository.CustomerRepository, parseService *service.ParseService) *CreateCustomerUseCase {
	return &CreateCustomerUseCase{
		repo:         repo,
		parseService: parseService,
	}
}

func (uc *CreateCustomerUseCase) Execute(input dto.InputCreateCustomerDto) (dto.OutputCreateCustomerDto, error) {
	customerDTO, err := uc.parseService.ExecuteParseService(input)

	if err != nil {
		return dto.OutputCreateCustomerDto{}, err
	}

	customer, err := entity.NewCustomer(
		customerDTO.Cpf,
		customerDTO.Private,
		customerDTO.Incompleto,
		customerDTO.DataUltimaCompra,
		customerDTO.TicketMedio,
		customerDTO.TicketUltimaCompra,
		customerDTO.LojaMaisFrequente,
		customerDTO.LojaUltimaCompra,
	)

	if err != nil {
		return dto.OutputCreateCustomerDto{}, err
	}

	err = uc.repo.Create(customer)

	if err != nil {
		return dto.OutputCreateCustomerDto{}, internalerrors.ErrInternal
	}

	output := dto.OutputCreateCustomerDto{
		ID:                          customer.ID,
		Cpf:                         customer.Cpf,
		CpfValido:                   customer.CpfValido,
		Private:                     customer.Private,
		Incompleto:                  customer.Incompleto,
		DataUltimaCompra:            customer.DataUltimaCompra,
		TicketMedio:                 customer.TicketMedio,
		TicketUltimaCompra:          customer.TicketUltimaCompra,
		LojaMaisFrequente:           customer.LojaMaisFrequente,
		CnpjLojaMaisFrequenteValido: customer.CnpjLojaMaisFrequenteValido,
		LojaUltimaCompra:            customer.LojaUltimaCompra,
		CnpjLojaUltimaCompraValido:  customer.CnpjLojaUltimaCompraValido,
		CreatedAt:                   customer.CreatedAt,
	}

	return output, nil
}
