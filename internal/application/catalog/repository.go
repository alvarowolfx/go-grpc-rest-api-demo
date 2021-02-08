package catalog

import (
	"context"

	tvtime "com.aviebrantz.tvtime/pkg/api"
)

type CatalogEntry struct {
	ID       string             `json:"id"`
	Name     string             `json:"name"`
	Slug     string             `json:"slug"`
	ImageURL string             `json:"imageUrl"`
	Type     tvtime.CatalogType `json:"type"`
}

type Repository interface {
	CreateCatalogEntry(ctx context.Context, param CreateCatalogEntryParam) (*CatalogEntry, error)
	FindCatalogEntryBySlug(ctx context.Context, slug string) (*CatalogEntry, error)
	ListEntries(ctx context.Context, filter EntriesFilter) ([]CatalogEntry, int, error)
}

type CreateCatalogEntryParam struct {
	Name     string              `json:"name" validate:"required"`
	Type     *tvtime.CatalogType `json:"type" `
	ImageURL string              `json:"imageUrl" validate:"required,url"`
}

// EntriesFilter parameters available to find/filter entries on the database
type EntriesFilter struct {
	ID      *string             `json:"id"`
	Term    *string             `json:"term"`
	Type    *tvtime.CatalogType `json:"type"`
	Page    int
	PerPage int
}
