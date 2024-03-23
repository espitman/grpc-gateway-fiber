package main

import (
	_ "github.com/espitman/grpc-gateway-fiber/docs"
	"github.com/gofiber/fiber/v2"
	"log"
)

// @title           gRPC Gateway
// @version         1.0
// @description     This is gRPC Gateway
// @contact.name   API Support
// @contact.email  s.heidar@jabama.com
// @BasePath  /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	app := fiber.New()

	priceServiceClient := newPriceServiceClient()
	priceServiceHandler := newPriceServiceHandler(priceServiceClient)

	router := newRouter(

		priceServiceHandler,
	)
	router.serve(app)
	log.Fatal(app.Listen(":8083"))
}
