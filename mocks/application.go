package mocks

import (
	"fase-4-hf-client/internal/core/domain/entity/dto"
	"strings"
)

type MockApplication struct {
	WantOut     *dto.OutputClient
	WantErr     error
	WantOutNull string
}

func (m MockApplication) GetClientByID(id string) (*dto.OutputClient, error) {
	if m.WantErr != nil && strings.EqualFold("errGetClientByID", m.WantErr.Error()) {
		return nil, m.WantErr
	}
	if strings.EqualFold(m.WantOutNull, "nilGetClientByID") {
		return nil, nil
	}
	return m.WantOut, nil
}

func (m MockApplication) GetClientByCPF(cpf string) (*dto.OutputClient, error) {
	if m.WantErr != nil && strings.EqualFold("errGetClientByCPF", m.WantErr.Error()) {
		return nil, m.WantErr
	}
	if strings.EqualFold(m.WantOutNull, "nilGetClientByCPF") {
		return nil, nil
	}
	return m.WantOut, nil
}

func (m MockApplication) SaveClient(reqClient dto.RequestClient) (*dto.OutputClient, error) {
	if m.WantErr != nil && strings.EqualFold("errSaveClient", m.WantErr.Error()) {
		return nil, m.WantErr
	}
	if strings.EqualFold(m.WantOutNull, "nilSaveClient") {
		return nil, nil
	}
	return m.WantOut, nil
}

// Repository Callers
type MockApplicationRepostoryCallers struct {
	WantOut *dto.ClientDB
	WantErr error
}

func (m MockApplicationRepostoryCallers) GetClientByCPFRepository(cpf string) (*dto.ClientDB, error) {
	if m.WantErr != nil && strings.EqualFold("errGetClientByCPFRepository", m.WantErr.Error()) {
		return nil, m.WantErr
	}
	return m.WantOut, nil
}

func (m MockApplicationRepostoryCallers) GetClientByIDRepository(cpf string) (*dto.ClientDB, error) {
	if m.WantErr != nil && strings.EqualFold("errGetClientByIDRepository", m.WantErr.Error()) {
		return nil, m.WantErr
	}
	return m.WantOut, nil
}

func (m MockApplicationRepostoryCallers) SaveClientRepository(client dto.ClientDB) (*dto.ClientDB, error) {
	if m.WantErr != nil && strings.EqualFold("errSaveClientRepository", m.WantErr.Error()) {
		return nil, m.WantErr
	}
	return m.WantOut, nil
}

// UseCase callers
type MockApplicationUseCaseCallers struct {
	WantErr error
}

func (m MockApplicationUseCaseCallers) GetClientByCPFUseCase(cpf string) error {
	if m.WantErr != nil && strings.EqualFold("errGetClientByCPFUseCase", m.WantErr.Error()) {
		return m.WantErr
	}
	return nil
}

func (m MockApplicationUseCaseCallers) GetClientByIDUseCase(cpf string) error {
	if m.WantErr != nil && strings.EqualFold("errGetClientByIDUseCase", m.WantErr.Error()) {
		return m.WantErr
	}
	return nil
}

func (m MockApplicationUseCaseCallers) SaveClientUseCase(client dto.RequestClient) error {
	if m.WantErr != nil && strings.EqualFold("errSaveClientUseCase", m.WantErr.Error()) {
		return m.WantErr
	}
	return nil
}
