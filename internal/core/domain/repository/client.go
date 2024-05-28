package repository

import (
	"fase-4-hf-client/internal/core/domain/entity/dto"
)

type ClientRepository interface {
	GetClientByID(uuid string) (*dto.ClientDB, error)
	GetClientByCPF(cpf string) (*dto.ClientDB, error)
	SaveClient(client dto.ClientDB) (*dto.ClientDB, error)
}
