package main

import (
	"context"
	"fmt"
	"log"
	"os"

	tvtime "com.aviebrantz.tvtime/pkg/api"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")
	grpcAddr := fmt.Sprintf(":%s", port)

	var conn *grpc.ClientConn
	conn, err = grpc.Dial(grpcAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	ctx := context.Background()
	client := tvtime.NewCatalogServiceClient(conn)
	_, err = client.Create(ctx, &tvtime.CreateEntryRequest{
		Name: "Wandavision",
		Type: tvtime.CatalogType_TV_SHOW,
	})
	if err != nil {
		log.Fatalf("failed to fetch running work orders: %v", err)
	}
}
