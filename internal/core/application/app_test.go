package application

import (
	"errors"
	ps "fase-4-hf-client/external/strings"
	"fase-4-hf-client/internal/core/domain/entity/dto"
	"fase-4-hf-client/mocks"
	"log"
	"testing"
)

// go test -v -count=1 -failfast -run ^Test_GetClientByCPF$
func Test_GetClientByCPF(t *testing.T) {
	type args struct {
		cpf string
	}

	tests := []struct {
		name           string
		args           args
		mockRepository mocks.MockClientRepository
		mockUseCase    mocks.MockClientUseCase
		wantOut        dto.OutputClient
		isWantedError  bool
	}{
		{
			name: "success",
			args: args{
				cpf: "10000000",
			},
			mockRepository: mocks.MockClientRepository{
				WantOut: &dto.ClientDB{
					UUID:      "001",
					Name:      "someUser",
					CPF:       "10000000",
					Email:     "someemail@some.com",
					CreatedAt: "2001-01-01 15:30:00",
				},
				WantErr: nil,
			},
			mockUseCase: mocks.MockClientUseCase{
				WantErr: nil,
			},
			wantOut: dto.OutputClient{
				UUID:      "001",
				Name:      "someUser",
				CPF:       "10000000",
				Email:     "someemail@some.com",
				CreatedAt: "2001-01-01 15:30:00",
			},
			isWantedError: false,
		},
		{
			name: "error_repository",
			args: args{
				cpf: "10000000",
			},
			mockRepository: mocks.MockClientRepository{
				WantOut: nil,
				WantErr: errors.New("errGetClientByCPF"),
			},
			mockUseCase: mocks.MockClientUseCase{
				WantErr: nil,
			},
			wantOut:       dto.OutputClient{},
			isWantedError: true,
		},
		{
			name: "error_useCase",
			args: args{
				cpf: "10000000",
			},
			mockRepository: mocks.MockClientRepository{
				WantOut: nil,
				WantErr: nil,
			},
			mockUseCase: mocks.MockClientUseCase{
				WantErr: errors.New("errGetClientByCPF"),
			},
			wantOut:       dto.OutputClient{},
			isWantedError: true,
		},
	}

	for _, tc := range tests {
		app := NewApplication(tc.mockRepository, tc.mockUseCase)
		t.Run(tc.name, func(*testing.T) {

			out, err := app.GetClientByCPF(tc.args.cpf)

			if (!tc.isWantedError) && err != nil {
				log.Panicf("unexpected error: %v", err)
			}

			if out != nil && (ps.MarshalString(out) != ps.MarshalString(tc.wantOut)) {
				t.Errorf("was not suppose to have:\n%s\n and got:\n%s\n", ps.MarshalString(tc.wantOut), ps.MarshalString(out))
			}
		})
	}
}

// go test -v -count=1 -failfast -run ^Test_SaveClient$
func Test_SaveClient(t *testing.T) {
	type args struct {
		req dto.RequestClient
	}

	tests := []struct {
		name           string
		args           args
		mockRepository mocks.MockClientRepository
		mockUseCase    mocks.MockClientUseCase
		wantOut        dto.OutputClient
		isWantedError  bool
	}{
		{
			name: "success",
			args: args{
				req: dto.RequestClient{
					UUID:      "001",
					Name:      "someUser",
					CPF:       "10000000",
					Email:     "someemail@some.com",
					CreatedAt: "2001-01-01 15:30:00",
				},
			},
			mockRepository: mocks.MockClientRepository{
				WantOutNull: "nilGetClientByCPF",
				WantOut: &dto.ClientDB{
					UUID:      "001",
					Name:      "someUser",
					CPF:       "10000000",
					Email:     "someemail@some.com",
					CreatedAt: "2001-01-01 15:30:00",
				},
				WantErr: nil,
			},
			mockUseCase: mocks.MockClientUseCase{
				WantErr: nil,
			},
			wantOut: dto.OutputClient{
				UUID:      "001",
				Name:      "someUser",
				CPF:       "10000000",
				Email:     "someemail@some.com",
				CreatedAt: "2001-01-01 15:30:00",
			},
			isWantedError: false,
		},
		{
			name: "success_null",
			args: args{
				req: dto.RequestClient{
					UUID:      "001",
					Name:      "someUser",
					CPF:       "10000000",
					Email:     "someemail@some.com",
					CreatedAt: "2001-01-01 15:30:00",
				},
			},
			mockRepository: mocks.MockClientRepository{
				WantOut: nil,
				WantErr: nil,
			},
			mockUseCase: mocks.MockClientUseCase{
				WantErr: nil,
			},
			wantOut: dto.OutputClient{
				UUID:      "001",
				Name:      "someUser",
				CPF:       "10000000",
				Email:     "someemail@some.com",
				CreatedAt: "2001-01-01 15:30:00",
			},
			isWantedError: false,
		},

		{
			name: "error_user_exists",
			args: args{
				req: dto.RequestClient{
					UUID:      "001",
					Name:      "someUser",
					CPF:       "10000000",
					Email:     "someemail@some.com",
					CreatedAt: "2001-01-01 15:30:00",
				},
			},
			mockRepository: mocks.MockClientRepository{
				WantOut: &dto.ClientDB{
					UUID:      "001",
					Name:      "someUser",
					CPF:       "10000000",
					Email:     "someemail@some.com",
					CreatedAt: "2001-01-01 15:30:00",
				},
				WantErr: nil,
			},
			mockUseCase: mocks.MockClientUseCase{
				WantErr: nil,
			},
			wantOut: dto.OutputClient{
				UUID:      "001",
				Name:      "someUser",
				CPF:       "10000000",
				Email:     "someemail@some.com",
				CreatedAt: "2001-01-01 15:30:00",
			},
			isWantedError: true,
		},
		{
			name: "error_useCase",
			args: args{
				req: dto.RequestClient{
					UUID:      "001",
					Name:      "someUser",
					CPF:       "10000000",
					Email:     "someemail@some.com",
					CreatedAt: "2001-01-01 15:30:00",
				},
			},
			mockRepository: mocks.MockClientRepository{
				WantOut: nil,
				WantErr: nil,
			},
			mockUseCase: mocks.MockClientUseCase{
				WantErr: errors.New("errSaveClient"),
			},
			wantOut: dto.OutputClient{
				UUID:      "001",
				Name:      "someUser",
				CPF:       "10000000",
				Email:     "someemail@some.com",
				CreatedAt: "2001-01-01 15:30:00",
			},
			isWantedError: true,
		},
		{
			name: "error_repository",
			args: args{
				req: dto.RequestClient{
					UUID:      "001",
					Name:      "someUser",
					CPF:       "10000000",
					Email:     "someemail@some.com",
					CreatedAt: "2001-01-01 15:30:00",
				},
			},
			mockRepository: mocks.MockClientRepository{
				WantOut: nil,
				WantErr: errors.New("errSaveClient"),
			},
			mockUseCase: mocks.MockClientUseCase{
				WantErr: nil,
			},
			wantOut: dto.OutputClient{
				UUID:      "001",
				Name:      "someUser",
				CPF:       "10000000",
				Email:     "someemail@some.com",
				CreatedAt: "2001-01-01 15:30:00",
			},
			isWantedError: true,
		},
		{
			name: "error_getUserByCPF",
			args: args{
				req: dto.RequestClient{
					UUID:      "001",
					Name:      "someUser",
					CPF:       "10000000",
					Email:     "someemail@some.com",
					CreatedAt: "2001-01-01 15:30:00",
				},
			},
			mockRepository: mocks.MockClientRepository{
				WantOut: nil,
				WantErr: errors.New("errGetClientByCPF"),
			},
			mockUseCase: mocks.MockClientUseCase{
				WantErr: nil,
			},
			wantOut: dto.OutputClient{
				UUID:      "001",
				Name:      "someUser",
				CPF:       "10000000",
				Email:     "someemail@some.com",
				CreatedAt: "2001-01-01 15:30:00",
			},
			isWantedError: true,
		},
	}

	for _, tc := range tests {
		app := NewApplication(tc.mockRepository, tc.mockUseCase)
		t.Run(tc.name, func(*testing.T) {
			out, err := app.SaveClient(tc.args.req)

			if (!tc.isWantedError) && err != nil {
				log.Panicf("unexpected error: %v", err)
			}

			if out != nil && (ps.MarshalString(out) != ps.MarshalString(tc.wantOut)) {
				t.Errorf("was not suppose to have:\n%s\n and got:\n%s\n", ps.MarshalString(tc.wantOut), ps.MarshalString(out))
			}
		})
	}
}
