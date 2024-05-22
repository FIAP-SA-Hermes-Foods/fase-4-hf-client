package application

import "fase-4-hf-client/internal/core/domain/entity/dto"

func (app application) GetClientByCPFUseCase(cpf string) error {
	return app.clientUC.GetClientByCPF(cpf)
}

func (app application) SaveClientUseCase(client dto.RequestClient) error {
	return app.clientUC.SaveClient(client)
}
