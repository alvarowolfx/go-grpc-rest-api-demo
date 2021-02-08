package favorites

import (
	"context"
)

type Repository interface {
	ListByUser(ctx context.Context, username string) ([]string, error)
	ListByMovie(ctx context.Context, catalogEntrySlug string) ([]string, error)
	AddFavorite(ctx context.Context, username, catalogEntrySlug string) error
	RemoveFavorite(ctx context.Context, username, catalogEntrySlug string) error
}
