package main

import (
	"context"
	"fmt"
	"log"

	"com.aviebrantz.tvtime/internal/application/catalog"
	tvtime "com.aviebrantz.tvtime/pkg/api"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/dgraph-io/badger/v3"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	db, err := badger.Open(badger.DefaultOptions("./tmp/badger"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	catalogRepo := catalog.NewBadgerCatalogRepository(db)
	catalogService := catalog.NewService(catalog.ServiceDeps{
		CatalogRepository: catalogRepo,
	})

	gofakeit.Seed(0)
	types := []int{
		int(tvtime.CatalogType_MOVIE),
		int(tvtime.CatalogType_TV_SHOW),
	}
	for i := 0; i < 1000; i++ {
		randType := tvtime.CatalogType(int32(gofakeit.RandomInt(types)))
		ctx := context.Background()
		entry, err := catalogService.CreateCatalogEntry(ctx, catalog.CreateCatalogEntryParam{
			Name:     gofakeit.AppName(),
			Type:     &randType,
			ImageURL: gofakeit.ImageURL(320, 240),
		})
		if err != nil {
			log.Fatalf("failed to create entry: %v", err)
		}
		fmt.Println(entry, randType)
	}
}
