package endpoints

import (
	"com.aviebrantz.tvtime/internal/application/catalog"
	tvtime "com.aviebrantz.tvtime/pkg/api"
)

func FromEntityToProtoCatalog(entry *catalog.CatalogEntry) *tvtime.CatalogEntry {
	return &tvtime.CatalogEntry{
		Id:       entry.ID,
		Slug:     entry.Slug,
		Name:     entry.Name,
		ImageUrl: entry.ImageURL,
		Type:     entry.Type,
	}
}
