package office

import (
	"context"
	"office/internal/repo/postgres"
	"office/internal/types"
	"office/pkg/office/pkg/prots"
)

type OfficeHandler struct {
	prots.UnimplementedOfficeServiceServer
	repo *postgres.Repo
}

func NewOffice(repo *postgres.Repo) *OfficeHandler {
	return &OfficeHandler{
		repo: repo,
	}
}

func (h *OfficeHandler) CreateOffice(ctx context.Context, req *prots.CreateOfficeRequest) (*prots.CreateOfficeResponse, error) {
	payload := types.OfficeMake{
		Name:    req.Name,
		Address: req.Address,
	}
	if err := h.repo.CreateOffice(ctx, payload); err != nil {
		return new(prots.CreateOfficeResponse), err
	}
	return new(prots.CreateOfficeResponse), nil
}

func (h *OfficeHandler) GetOfficeList(ctx context.Context, req *prots.GetOfficeListRequest) (*prots.GetOfficeListResponse, error) {
	_, err := h.repo.GetOffices(ctx)
	if err != nil {
		return new(prots.GetOfficeListResponse), err
	}
	return new(prots.GetOfficeListResponse), nil
}
