package entity

import (
	vo "fase-4-hf-client/internal/core/domain/entity/valueObject"
)

type Client struct {
	ID        int64        `json:"id,omitempty"`
	Name      string       `json:"name,omitempty"`
	CPF       vo.Cpf       `json:"cpf,omitempty"`
	Email     string       `json:"email,omitempty"`
	CreatedAt vo.CreatedAt `json:"createdAt,omitempty"`
}
