package useCase

import (
	"fase-4-hf-client/internal/core/domain/entity/dto"
	"log"
	"testing"
)

// go test -v -failfast -run ^Test_GetClientByCPF$
func Test_GetClientByCPF(t *testing.T) {
	type args struct {
		cpf string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				cpf: "1000000000",
			},
			wantErr: false,
		},
		{
			name: "not_valid_cpf",
			args: args{
				cpf: "",
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		uc := NewClientUseCase()
		t.Run(tc.name, func(*testing.T) {
			err := uc.GetClientByCPF(tc.args.cpf)
			if (!tc.wantErr) && err != nil {
				log.Panicf("unexpected error: %v", err)
			}
		})
	}
}

// go test -v -failfast -run ^Test_SaveClient$
func Test_SaveClient(t *testing.T) {

	type args struct {
		reqClient dto.RequestClient
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				reqClient: dto.RequestClient{
					UUID:      "",
					Name:      "",
					CPF:       "10000000",
					Email:     "",
					CreatedAt: "",
				},
			},
			wantErr: false,
		},
		{
			name: "not_valid_cpf",
			args: args{
				reqClient: dto.RequestClient{
					UUID:      "",
					Name:      "",
					CPF:       "",
					Email:     "",
					CreatedAt: "",
				},
			},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		uc := NewClientUseCase()
		t.Run(tc.name, func(*testing.T) {
			err := uc.SaveClient(tc.args.reqClient)
			if (!tc.wantErr) && err != nil {
				log.Panicf("unexpected error: %v", err)
			}
		})
	}

}
