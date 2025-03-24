package usecase

import (
	"neoway_test/internal/domain/customer/dto"
	"neoway_test/internal/domain/customer/repository"
	internalerrors "neoway_test/internal/internal-errors"
)

type GetCustomerByCpfUseCase struct {
	repo repository.CustomerRepository
}

func NewGetCustomerByCpfUseCase(repo repository.CustomerRepository) *GetCustomerByCpfUseCase {
	return &GetCustomerByCpfUseCase{repo: repo}
}

func (uc *GetCustomerByCpfUseCase) Execute(input dto.InputGetCustomerByCpfDto) (*dto.OutputGetCustomerDto, error) {
	customer, err := uc.repo.GetByCpf(input.Cpf)

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
