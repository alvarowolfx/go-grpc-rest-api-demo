package favorites

import (
	"context"
	"errors"

	"com.aviebrantz.tvtime/internal/application/catalog"
	"com.aviebrantz.tvtime/internal/application/users"
)

type Service interface {
	Repository
}

type ServiceDeps struct {
	FavoriteRepository Repository
	UserRepository     users.Repository
	CatalogRepository  catalog.Repository
}

type favoriteService struct {
	ServiceDeps
}

func NewService(deps ServiceDeps) Service {
	return &favoriteService{deps}
}

func (s *favoriteService) ListByUser(ctx context.Context, username string) ([]string, error) {
	user, err := s.UserRepository.FindUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	return s.FavoriteRepository.ListByUser(ctx, username)
}

func (s *favoriteService) ListByMovie(ctx context.Context, catalogEntrySlug string) ([]string, error) {
	entry, err := s.CatalogRepository.FindCatalogEntryBySlug(ctx, catalogEntrySlug)
	if err != nil {
		return nil, err
	}

	if entry == nil {
		return nil, errors.New("catalog entry not found")
	}

	return s.FavoriteRepository.ListByMovie(ctx, catalogEntrySlug)
}

func (s *favoriteService) AddFavorite(ctx context.Context, username, catalogEntrySlug string) error {
	favorites, err := s.ListByUser(ctx, username)
	if err != nil {
		return err
	}
	for _, f := range favorites {
		if f == catalogEntrySlug {
			return errors.New("already favorited")
		}
	}

	entry, err := s.CatalogRepository.FindCatalogEntryBySlug(ctx, catalogEntrySlug)
	if err != nil {
		return err
	}

	if entry == nil {
		return errors.New("catalog entry not found")
	}

	return s.FavoriteRepository.AddFavorite(ctx, username, catalogEntrySlug)
}

func (s *favoriteService) RemoveFavorite(ctx context.Context, username, catalogEntrySlug string) error {
	user, err := s.UserRepository.FindUserByUsername(ctx, username)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("user not found")
	}

	entry, err := s.CatalogRepository.FindCatalogEntryBySlug(ctx, catalogEntrySlug)
	if err != nil {
		return err
	}

	if entry == nil {
		return errors.New("catalog entry not found")
	}

	return s.FavoriteRepository.RemoveFavorite(ctx, username, catalogEntrySlug)
}
