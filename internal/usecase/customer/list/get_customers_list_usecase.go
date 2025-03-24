package usecase

import (
	"neoway_test/internal/domain/customer/dto"
	"neoway_test/internal/domain/customer/repository"
	internalerrors "neoway_test/internal/internal-errors"
)

type GetCustomersListUseCase struct {
	repo repository.CustomerRepository
}

func NewGetCustomersListUseCase(repo repository.CustomerRepository) *GetCustomersListUseCase {
	return &GetCustomersListUseCase{repo: repo}
}

func (uc *GetCustomersListUseCase) Execute(input dto.InputGetCustomersListDto) ([]*dto.OutputGetCustomersListDto, error) {
	customerList, err := uc.repo.Get(input.Page)

	if err != nil {
		return nil, internalerrors.ProcessErrorToReturn(err)
	}

	var customersDto []*dto.OutputGetCustomersListDto
	for _, customer := range customerList {
		filaLojaDto := &dto.OutputGetCustomersListDto{
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
		customersDto = append(customersDto, filaLojaDto)
	}

	return customersDto, nil
}
