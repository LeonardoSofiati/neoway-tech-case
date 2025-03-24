package usecase

import (
	"neoway_test/internal/domain/customer/dto"
	"neoway_test/internal/domain/customer/repository"
	internalerrors "neoway_test/internal/internal-errors"
)

type GetCustomerByIdUseCase struct {
	repo repository.CustomerRepository
}

func NewGetCustomerByIdUseCase(repo repository.CustomerRepository) *GetCustomerByIdUseCase {
	return &GetCustomerByIdUseCase{repo: repo}
}

func (uc *GetCustomerByIdUseCase) Execute(input dto.InputGetCustomerByIdDto) (*dto.OutputGetCustomerDto, error) {
	customer, err := uc.repo.GetById(input.ID)

	if err != nil {
		return nil, internalerrors.ProcessErrorToReturn(err)
	}

	return &dto.OutputGetCustomerDto{
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
	}, nil
}
