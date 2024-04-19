package application

import (
	"errors"
	l "fase-4-hf-client/external/logger"
	ps "fase-4-hf-client/external/strings"
	"fase-4-hf-client/internal/core/domain/entity/dto"
	"fase-4-hf-client/internal/core/domain/repository"
	"fase-4-hf-client/internal/core/domain/useCase"

	"github.com/google/uuid"
)

type Application interface {
	GetClientByCPF(cpf string) (*dto.OutputClient, error)
	SaveClient(reqClient dto.RequestClient) (*dto.OutputClient, error)
}

type application struct {
	clientRepo repository.ClientRepository
	clientUC   useCase.ClientUseCase
}

func NewApplication(clientRepo repository.ClientRepository, clientUC useCase.ClientUseCase) Application {
	return application{clientRepo: clientRepo, clientUC: clientUC}
}

func (app application) GetClientByCPF(cpf string) (*dto.OutputClient, error) {
	l.Infof("GetClientByCPFApp: ", " | ", cpf)
	if err := app.GetClientByCPFUseCase(cpf); err != nil {
		l.Errorf("GetClientByCPFApp error: ", " | ", err)
		return nil, err
	}

	cOutDb, err := app.GetClientByCPFRepository(cpf)

	if err != nil {
		l.Errorf("GetClientByCPFApp error: ", " | ", err)
		return nil, err
	}

	if cOutDb == nil {
		l.Infof("GetClientByCPFApp output: ", " | ", cOutDb)
		return nil, nil
	}

	out := &dto.OutputClient{
		UUID:      cOutDb.UUID,
		Name:      cOutDb.Name,
		CPF:       cOutDb.CPF,
		Email:     cOutDb.Email,
		CreatedAt: cOutDb.CreatedAt,
	}

	l.Infof("GetClientByCPFApp output: ", " | ", ps.MarshalString(out))
	return out, err
}

func (app application) SaveClient(client dto.RequestClient) (*dto.OutputClient, error) {
	l.Infof("SaveClientApp: ", " | ", ps.MarshalString(client))
	clientWithCpf, err := app.GetClientByCPF(client.CPF)

	if err != nil {
		l.Errorf("SaveClientApp error: ", " | ", err)
		return nil, err
	}

	if clientWithCpf != nil {
		l.Errorf("SaveClientApp error: ", " | ", "is not possible to save client because this cpf is already in use")
		return nil, errors.New("is not possible to save client because this cpf is already in use")
	}

	if err := app.SaveClientUseCase(client); err != nil {
		l.Errorf("SaveClientApp error: ", " | ", err)
		return nil, err
	}

	clientDB := dto.ClientDB{
		UUID:      uuid.NewString(),
		Name:      client.Name,
		CPF:       client.Name,
		Email:     client.Email,
		CreatedAt: client.CreatedAt,
	}

	cOutDb, err := app.SaveClientRepository(clientDB)

	if err != nil {
		l.Errorf("SaveClientApp error: ", " | ", err)
		return nil, err
	}

	if cOutDb == nil {
		l.Infof("SaveClientApp output: ", " | ", nil)
		return nil, nil
	}

	out := &dto.OutputClient{
		UUID:      cOutDb.UUID,
		Name:      cOutDb.Name,
		CPF:       cOutDb.CPF,
		Email:     cOutDb.Email,
		CreatedAt: cOutDb.CreatedAt,
	}

	l.Infof("SaveClientApp output: ", " | ", ps.MarshalString(out))

	return out, nil
}
