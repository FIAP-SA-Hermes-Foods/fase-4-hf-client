package rpc

import (
	"context"
	cp "fase-4-hf-client/client_proto"
	"fase-4-hf-client/internal/core/application"
	"fase-4-hf-client/internal/core/domain/entity/dto"
)

type HandlerGRPC interface {
	Handler() *handlerGRPC
}

type handlerGRPC struct {
	app application.Application
	cp.UnimplementedClientServer
}

func NewHandler(app application.Application) HandlerGRPC {
	return &handlerGRPC{app: app}
}

func (h *handlerGRPC) Handler() *handlerGRPC {
	return h
}

func (h *handlerGRPC) CreateClient(ctx context.Context, req *cp.CreateClientRequest) (*cp.CreateClientResponse, error) {

	input := dto.RequestClient{
		Name:  req.Name,
		CPF:   req.Cpf,
		Email: req.Email,
	}

	c, err := h.app.SaveClient(input)

	if err != nil {
		return nil, err
	}

	if c == nil {
		return nil, nil
	}

	out := &cp.CreateClientResponse{
		Uuid:      c.UUID,
		Name:      c.Name,
		Cpf:       c.CPF,
		Email:     c.Email,
		CreatedAt: c.CreatedAt,
	}

	return out, nil

}

func (h *handlerGRPC) GetClientByCPF(ctx context.Context, req *cp.GetClientByCPFRequest) (*cp.GetClientByCPFResponse, error) {
	c, err := h.app.GetClientByCPF(req.Cpf)

	if err != nil {
		return nil, err
	}

	if c == nil {
		return nil, nil
	}

	out := &cp.GetClientByCPFResponse{
		Uuid:      c.UUID,
		Name:      c.Name,
		Cpf:       c.CPF,
		Email:     c.Email,
		CreatedAt: c.CreatedAt,
	}

	return out, nil
}
