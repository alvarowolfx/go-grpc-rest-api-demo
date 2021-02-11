package server

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"strings"

	"com.aviebrantz.tvtime/internal/application/auth"
	"com.aviebrantz.tvtime/internal/application/catalog"
	"com.aviebrantz.tvtime/internal/application/favorites"
	"com.aviebrantz.tvtime/internal/application/users"
	"com.aviebrantz.tvtime/internal/server/endpoints"
	"com.aviebrantz.tvtime/internal/server/middleware"
	tvtime "com.aviebrantz.tvtime/pkg/api"
	"github.com/dgraph-io/badger/v3"
	"github.com/gobuffalo/packr/v2"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rakyll/statik/fs"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"

	// Static files
	_ "com.aviebrantz.tvtime/statik"
)

// grpcHandlerFunc returns an http.Handler that delegates to grpcServer on incoming gRPC
// connections or otherHandler otherwise. Copied from cockroachdb.
func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}

func getOpenAPIHandler() http.Handler {
	_ = mime.AddExtensionType(".svg", "image/svg+xml")

	statikFS, err := fs.New()
	if err != nil {
		// Panic since this is a permanent error.
		log.Fatalf("creating OpenAPI filesystem: " + err.Error())
	}

	return http.FileServer(statikFS)
}

func Run() error {

	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)

	db, err := badger.Open(badger.DefaultOptions("./tmp/badger"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	usersRepo := users.NewBadgerUserRepository(db)
	favoritesRepo := favorites.NewBadgerFavoriteRepository(db)
	catalogRepo := catalog.NewBadgerCatalogRepository(db)
	userService := users.NewService(users.ServiceDeps{
		UserRepository: usersRepo,
	})
	authService := auth.NewService(auth.ServiceDeps{
		UserService: userService,
		JWTSecret:   os.Getenv("JWT_SECRET"),
	})
	catalogService := catalog.NewService(catalog.ServiceDeps{
		CatalogRepository: catalogRepo,
	})
	favoriteService := favorites.NewService(favorites.ServiceDeps{
		FavoriteRepository: favoritesRepo,
		CatalogRepository:  catalogRepo,
		UserRepository:     usersRepo,
	})

	grpcServerOpts := []grpc.ServerOption{}
	authInterceptor := middleware.AuthUnaryInterceptor(authService)
	grpcServerOpts = append(
		grpcServerOpts,
		grpc.UnaryInterceptor(authInterceptor),
	)

	authEndpoint := endpoints.NewAuthEndpoint(endpoints.AuthEndpointDeps{
		AuthService: authService,
	})

	userEndpoint := endpoints.NewUsersEndpoint(endpoints.UsersEndpointDeps{
		UserService:     userService,
		FavoriteService: favoriteService,
		CatalogService:  catalogService,
	})

	catalogEndpoint := endpoints.NewCatalogEndpoint(endpoints.CatalogEndpointDeps{
		CatalogService:  catalogService,
		FavoriteService: favoriteService,
	})

	grpcServer := grpc.NewServer(grpcServerOpts...)
	tvtime.RegisterAuthServiceServer(grpcServer, authEndpoint)
	tvtime.RegisterUserServiceServer(grpcServer, userEndpoint)
	tvtime.RegisterCatalogServiceServer(grpcServer, catalogEndpoint)

	// Start HTTP server that serves both Rest Gateway and GRPC
	mux := http.NewServeMux()

	// Register gRPC server endpoint
	gwmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	ctx := context.Background()
	err = tvtime.RegisterAuthServiceHandlerFromEndpoint(ctx, gwmux, addr, opts)
	if err != nil {
		log.Fatalf("Failed to start grpc server: %v\n", err)
		return err
	}
	err = tvtime.RegisterCatalogServiceHandlerFromEndpoint(ctx, gwmux, addr, opts)
	if err != nil {
		log.Fatalf("Failed to start grpc server: %v\n", err)
		return err
	}
	err = tvtime.RegisterUserServiceHandlerFromEndpoint(ctx, gwmux, addr, opts)
	if err != nil {
		log.Fatalf("Failed to start grpc server: %v\n", err)
		return err
	}

	mux.Handle("/static/", http.StripPrefix("/static/", getOpenAPIHandler()))
	mux.Handle("/", gwmux)
	swaggerBox := packr.New("swagger.json", "../../api")
	mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		swagger, err := swaggerBox.FindString("tvtime.swagger.json")
		if err != nil {
			w.WriteHeader(404)
			io.WriteString(w, "Swagger file not found")
			return
		}
		io.Copy(w, strings.NewReader(swagger))
	})
	muxWithCors := allowCORS(mux)

	return http.ListenAndServe(
		addr,
		grpcHandlerFunc(grpcServer, muxWithCors),
	)
}
