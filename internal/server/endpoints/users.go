package endpoints

import (
	"context"

	"com.aviebrantz.tvtime/internal/application/catalog"
	"com.aviebrantz.tvtime/internal/application/favorites"
	"com.aviebrantz.tvtime/internal/application/users"
	"com.aviebrantz.tvtime/internal/server/middleware"
	tvtime "com.aviebrantz.tvtime/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserEndpoint interface {
	tvtime.UserServiceServer
}

type UsersEndpointDeps struct {
	UserService     users.Service
	FavoriteService favorites.Service
	CatalogService  catalog.Service
}

type userEndpoint struct {
	UsersEndpointDeps
}

// NewUsersEndpoint
func NewUsersEndpoint(deps UsersEndpointDeps) UserEndpoint {
	return &userEndpoint{deps}
}

func (e *userEndpoint) RegisterUser(ctx context.Context, req *tvtime.RegisterUserRequest) (*tvtime.RegisterUserResponse, error) {
	user, err := e.UserService.CreateUser(ctx, users.CreateUserParam{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to create user: %v", err)
	}
	return &tvtime.RegisterUserResponse{
		User: &tvtime.User{
			Id:       user.ID,
			Username: user.Username,
		},
	}, nil
}

func (e *userEndpoint) ListFavorites(ctx context.Context, req *tvtime.ListFavoritesRequest) (*tvtime.ListFavoritesResponse, error) {
	userID := middleware.GetUserIDFromContext(ctx)
	catalogEntries, err := e.FavoriteService.ListByUser(ctx, userID)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to list favorites: %v", err)
	}
	entries := []*tvtime.CatalogEntry{}
	for _, c := range catalogEntries {
		entry, err := e.CatalogService.FindCatalogEntryBySlug(ctx, c)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to fetch favorited movie: %v", err)
		}
		entries = append(entries, FromEntityToProtoCatalog(entry))
	}
	return &tvtime.ListFavoritesResponse{

		Entries: entries,
	}, nil
}

func (e *userEndpoint) AddMovieToFavorites(ctx context.Context, req *tvtime.AddRemoveMovieFavoriteRequest) (*tvtime.AddRemoveMovieFavoriteResponse, error) {
	userID := middleware.GetUserIDFromContext(ctx)
	err := e.FavoriteService.AddFavorite(ctx, userID, req.MovieId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed add movie to favorite: %v", err)
	}
	return &tvtime.AddRemoveMovieFavoriteResponse{}, nil
}

func (e *userEndpoint) RemoveMovieFromFavorites(ctx context.Context, req *tvtime.AddRemoveMovieFavoriteRequest) (*tvtime.AddRemoveMovieFavoriteResponse, error) {
	userID := middleware.GetUserIDFromContext(ctx)
	err := e.FavoriteService.RemoveFavorite(ctx, userID, req.MovieId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to remove movie from favorite: %v", err)
	}
	return &tvtime.AddRemoveMovieFavoriteResponse{}, nil
}
