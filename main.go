package main

import (
	"shop/handler"
	pb "shop/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("shop"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterShopHandler(srv.Server(), new(handler.Shop))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
