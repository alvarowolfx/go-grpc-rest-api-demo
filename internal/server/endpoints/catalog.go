package endpoints

import (
	"context"

	"com.aviebrantz.tvtime/internal/application/catalog"
	"com.aviebrantz.tvtime/internal/application/favorites"
	tvtime "com.aviebrantz.tvtime/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CatalogEndpoint interface {
	tvtime.CatalogServiceServer
}

type CatalogEndpointDeps struct {
	CatalogService  catalog.Service
	FavoriteService favorites.Service
}

type catalogEndpoint struct {
	CatalogEndpointDeps
}

// NewCatalogEndpoint
func NewCatalogEndpoint(deps CatalogEndpointDeps) CatalogEndpoint {
	return &catalogEndpoint{deps}
}

func (e *catalogEndpoint) Create(ctx context.Context, req *tvtime.CreateEntryRequest) (*tvtime.CreateEntryResponse, error) {
	entry, err := e.CatalogService.CreateCatalogEntry(ctx, catalog.CreateCatalogEntryParam{
		Name:     req.Name,
		Type:     &req.Type,
		ImageURL: req.ImageUrl,
	})

	if err != nil {
		if _, ok := status.FromError(err); ok {
			return nil, err
		}
		return nil, status.Errorf(codes.FailedPrecondition, "failed to create entry: %v", err)
	}

	return &tvtime.CreateEntryResponse{
		Entry: FromEntityToProtoCatalog(entry),
	}, nil
}

func (e *catalogEndpoint) List(ctx context.Context, req *tvtime.SearchCatalogFilter) (*tvtime.ListResponse, error) {
	if req.PerPage <= 0 {
		req.PerPage = 10
	}
	if req.Page <= 0 {
		req.Page = 0
	}
	var term *string
	if req.Term != nil {
		term = &req.Term.Value
	}
	entries, total, err := e.CatalogService.ListEntries(ctx, catalog.EntriesFilter{
		Term:    term,
		Type:    &req.Type,
		Page:    int(req.Page),
		PerPage: int(req.PerPage),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list entry: %v", err)
	}
	response := &tvtime.ListResponse{
		Entries: []*tvtime.CatalogEntry{},
		Total:   int64(total),
		Page:    int32(req.Page),
		PerPage: int32(req.PerPage),
	}
	for _, entry := range entries {
		response.Entries = append(response.Entries, FromEntityToProtoCatalog(&entry))
	}
	return response, nil
}

func (e *catalogEndpoint) Get(ctx context.Context, req *tvtime.GetCatalogRequest) (*tvtime.GetCatalogResponse, error) {
	entry, err := e.CatalogService.FindCatalogEntryBySlug(ctx, req.Slug)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "failed find entry: %v", err)
	}
	favorites, err := e.FavoriteService.ListByMovie(ctx, req.Slug)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed get favourites entry: %v", err)
	}
	return &tvtime.GetCatalogResponse{
		Entry:     FromEntityToProtoCatalog(entry),
		Favorited: int64(len(favorites)),
	}, nil
}
