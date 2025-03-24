package usecase

import (
	"neoway_test/internal/domain/customer/dto"
	"neoway_test/internal/domain/customer/repository"
	internalerrors "neoway_test/internal/internal-errors"
)

type DeleteCustomerUseCase struct {
	repo repository.CustomerRepository
}

func NewDeleteCustomerUseCase(repo repository.CustomerRepository) *DeleteCustomerUseCase {
	return &DeleteCustomerUseCase{repo: repo}
}

func (uc *DeleteCustomerUseCase) Execute(input dto.InputDeleteCustomerDto) error {
	customerFound, err := uc.repo.GetById(input.ID)

	if err != nil {
		return err
	}

	err = uc.repo.Delete(customerFound)
	if err != nil {
		return internalerrors.ErrInternal
	}

	return nil
}
