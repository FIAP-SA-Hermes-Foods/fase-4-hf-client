package bdd

import (
	"context"
	"fase-4-hf-client/external/db/dynamo"
	l "fase-4-hf-client/external/logger"
	ps "fase-4-hf-client/external/strings"
	repositories "fase-4-hf-client/internal/adapters/driven/repositories/nosql"
	"fase-4-hf-client/internal/core/application"
	"fase-4-hf-client/internal/core/domain/entity/dto"
	"fase-4-hf-client/internal/core/useCase"
	"fmt"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/marcos-dev88/genv"
)

// go test -v -count=1 -failfast -run ^Test_GetClientByCPF$
func Test_GetClientByCPF(t *testing.T) {
	if err := genv.New("../../.env"); err != nil {
		l.Errorf("error set envs %v", " | ", err)
	}

	l.Info("====> TEST GetClientByCPF <====")

	type Input struct {
		CPF string `json:"cpf"`
	}

	type Output struct {
		Output *dto.OutputClient `json:"output"`
	}

	tests := []struct {
		scenario          string
		name              string
		input             Input
		shouldReturnError bool
		shouldBeNull      bool
		expectedOutput    Output
	}{
		{
			scenario: "Sending a valid and existing CPF",
			name:     "success_valid_cpf",
			input: Input{
				CPF: "12345678901",
			},
			shouldReturnError: false,
			shouldBeNull:      false,
			expectedOutput: Output{
				Output: &dto.OutputClient{
					UUID:      "1",
					Name:      "leo",
					CPF:       "12345678901",
					Email:     "leo@some.com",
					CreatedAt: "19-05-2024 23:13:29",
				},
			},
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(*testing.T) {

			// config
			ctx := context.Background()

			cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-east-1"))
			if err != nil {
				log.Fatalf("unable to load SDK config, %v", err)
			}

			db := dynamo.NewDynamoDB(cfg)
			repo := repositories.NewClientRepository(db, os.Getenv("DB_TABLE"))
			uc := useCase.NewClientUseCase()
			app := application.NewApplication(repo, uc)
			// final config

			l.Info("----------------")
			l.Info(fmt.Sprintf("-> Scenario: %s", tc.scenario))
			l.Info(fmt.Sprintf("-> Input: %s", ps.MarshalString(tc.input)))
			l.Info(fmt.Sprintf("-> ExpectedOutput: %s", ps.MarshalString(tc.expectedOutput)))
			l.Info(fmt.Sprintf("-> Should return error: %v", tc.shouldReturnError))
			l.Info("----------------")

			client, err := app.GetClientByCPF(tc.input.CPF)

			if (!tc.shouldReturnError) && err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if !tc.shouldBeNull {
				if client.Name != tc.expectedOutput.Output.Name {
					t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output.Name, client.Name)
				}

				if client.CPF != tc.expectedOutput.Output.CPF {
					t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output.CPF, client.CPF)
				}

				if client.Email != tc.expectedOutput.Output.Email {
					t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output.Email, client.Email)
				}

			}
			l.Info(fmt.Sprintf("====> Success running scenario: [%s] <====", tc.scenario))
		})
		l.Info("====> Finish BDD Test GetClientByCPF <====")
	}
}

// go test -v -count=1 -failfast -run ^Test_SaveClient$
func Test_SaveClient(t *testing.T) {
	if err := genv.New("../../.env"); err != nil {
		l.Errorf("error set envs %v", " | ", err)
	}

	const chars = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	generateRandomChar := func(length int, charset string) string {
		b := make([]byte, length)
		for i := range b {
			b[i] = charset[seededRand.Intn(len(charset))]
		}
		return string(b)
	}

	cpfRandomValue := generateRandomChar(11, chars)

	l.Info("====> TEST SaveClient <====")

	type Input struct {
		Input *dto.RequestClient `json:"input"`
	}

	type Output struct {
		Output *dto.OutputClient `json:"output"`
	}

	tests := []struct {
		scenario          string
		name              string
		input             Input
		shouldReturnError bool
		shouldBeNull      bool
		expectedOutput    Output
	}{
		{
			scenario: "Sending a valid input",
			name:     "success",
			input: Input{
				Input: &dto.RequestClient{
					Name:      "Marty",
					CPF:       cpfRandomValue,
					Email:     "marty_flying@email.com",
					CreatedAt: "19-05-2024 23:13:29",
				},
			},
			shouldReturnError: false,
			shouldBeNull:      false,
			expectedOutput: Output{
				Output: &dto.OutputClient{
					Name:      "Marty",
					CPF:       cpfRandomValue,
					Email:     "marty_flying@email.com",
					CreatedAt: "19-05-2024 23:13:29",
				},
			},
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(*testing.T) {

			// config
			ctx := context.Background()

			cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-east-1"))
			if err != nil {
				log.Fatalf("unable to load SDK config, %v", err)
			}

			db := dynamo.NewDynamoDB(cfg)
			repo := repositories.NewClientRepository(db, os.Getenv("DB_TABLE"))
			uc := useCase.NewClientUseCase()
			app := application.NewApplication(repo, uc)
			// final config

			l.Info("----------------")
			l.Info(fmt.Sprintf("-> Scenario: %s", tc.scenario))
			l.Info(fmt.Sprintf("-> Input: %s", ps.MarshalString(tc.input)))
			l.Info(fmt.Sprintf("-> ExpectedOutput: %s", ps.MarshalString(tc.expectedOutput)))
			l.Info(fmt.Sprintf("-> Should return error: %v", tc.shouldReturnError))
			l.Info("----------------")

			client, err := app.SaveClient(*tc.input.Input)

			if (!tc.shouldReturnError) && err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if !tc.shouldBeNull {
				if client.Name != tc.expectedOutput.Output.Name {
					t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output.Name, client.Name)
				}

				if client.CPF != tc.expectedOutput.Output.CPF {
					t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output.CPF, client.CPF)
				}

				if client.Email != tc.expectedOutput.Output.Email {
					t.Errorf("expected: %s\ngot: %s", tc.expectedOutput.Output.Email, client.Email)
				}

			}
			l.Info(fmt.Sprintf("====> Success running scenario: [%s] <====", tc.scenario))
		})
		l.Info("====> Finish BDD Test SaveClient <====")
	}
}
