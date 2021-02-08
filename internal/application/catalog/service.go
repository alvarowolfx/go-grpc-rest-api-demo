package catalog

import (
	"context"

	"com.aviebrantz.tvtime/pkg/validations"
)

type Service interface {
	Repository
}

type ServiceDeps struct {
	CatalogRepository Repository
}

type catalogService struct {
	ServiceDeps
}

func NewService(deps ServiceDeps) Service {
	return &catalogService{deps}
}

func (s *catalogService) CreateCatalogEntry(ctx context.Context, param CreateCatalogEntryParam) (*CatalogEntry, error) {
	errs := validations.ValidateStruct(param)
	if len(errs) != 0 {
		return nil, validations.GroupValidationErrorAsGrpcError("failed to create entry", errs)
	}
	return s.CatalogRepository.CreateCatalogEntry(ctx, param)
}

func (s *catalogService) ListEntries(ctx context.Context, filter EntriesFilter) ([]CatalogEntry, int, error) {
	return s.CatalogRepository.ListEntries(ctx, filter)
}

func (s *catalogService) FindCatalogEntryBySlug(ctx context.Context, slug string) (*CatalogEntry, error) {
	return s.CatalogRepository.FindCatalogEntryBySlug(ctx, slug)
}
