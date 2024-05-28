package mocks

import (
	"fase-4-hf-client/internal/core/domain/entity/dto"
	"strings"
)

type MockClientUseCase struct {
	WantOutNull string
	WantErr     error
}

func (m MockClientUseCase) GetClientByID(cpf string) error {
	if m.WantErr != nil && strings.EqualFold("errGetClientByID", m.WantErr.Error()) {
		return m.WantErr
	}
	return nil
}

func (m MockClientUseCase) GetClientByCPF(cpf string) error {
	if m.WantErr != nil && strings.EqualFold("errGetClientByCPF", m.WantErr.Error()) {
		return m.WantErr
	}
	return nil
}

func (m MockClientUseCase) SaveClient(reqClient dto.RequestClient) error {
	if m.WantErr != nil && strings.EqualFold("errSaveClient", m.WantErr.Error()) {
		return m.WantErr
	}

	return nil
}
