package useCase

import "fase-4-hf-client/internal/core/domain/entity/dto"

type ClientUseCase interface {
	GetClientByID(uuid string) error
	GetClientByCPF(cpf string) error
	SaveClient(reqClient dto.RequestClient) error
}
