package useCase

import (
	"errors"
	"fase-4-hf-client/internal/core/domain/entity/dto"
	"fase-4-hf-client/internal/core/domain/useCase"
)

var _ useCase.ClientUseCase = (*clientUseCase)(nil)

type clientUseCase struct {
}

func NewClientUseCase() clientUseCase {
	return clientUseCase{}
}

func (c clientUseCase) SaveClient(reqClient dto.RequestClient) error {
	client := reqClient.Client()

	if err := client.CPF.Validate(); err != nil {
		return err
	}

	return nil
}

func (c clientUseCase) GetClientByCPF(cpf string) error {
	if len(cpf) == 0 {
		return errors.New("the cpf is not valid for consult")
	}

	return nil
}

func (c clientUseCase) GetClientByID(uuid string) error {
	if len(uuid) == 0 {
		return errors.New("the uuid is not valid for consult")
	}
	return nil
}
