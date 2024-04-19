package useCase

import "fase-4-hf-client/internal/core/domain/entity/dto"

type ClientUseCase interface {
	SaveClient(reqClient dto.RequestClient) error
	GetClientByCPF(cpf string) error
}
