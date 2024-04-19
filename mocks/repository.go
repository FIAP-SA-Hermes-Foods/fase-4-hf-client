package mocks

import (
	"fase-4-hf-client/internal/core/domain/entity/dto"
	"strings"
)

type MockClientRepository struct {
	WantOut     *dto.ClientDB
	WantOutNull string
	WantErr     error
}

func (m MockClientRepository) GetClientByCPF(cpf string) (*dto.ClientDB, error) {
	if m.WantErr != nil && strings.EqualFold("errGetClientByCPF", m.WantErr.Error()) {
		return nil, m.WantErr
	}
	if strings.EqualFold(m.WantOutNull, "nilGetClientByCPF") {
		return nil, nil
	}
	return m.WantOut, nil
}

func (m MockClientRepository) SaveClient(client dto.ClientDB) (*dto.ClientDB, error) {
	if m.WantErr != nil && strings.EqualFold("errSaveClient", m.WantErr.Error()) {
		return nil, m.WantErr
	}

	if strings.EqualFold(m.WantOutNull, "nilSaveClient") {
		return nil, nil
	}

	return m.WantOut, nil
}
