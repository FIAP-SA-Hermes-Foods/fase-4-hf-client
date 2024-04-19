package application

import "fase-4-hf-client/internal/core/domain/entity/dto"

func (app application) GetClientByCPFRepository(cpf string) (*dto.ClientDB, error) {
	return app.clientRepo.GetClientByCPF(cpf)
}

func (app application) SaveClientRepository(client dto.ClientDB) (*dto.ClientDB, error) {
	return app.clientRepo.SaveClient(client)
}
