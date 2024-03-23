package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

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
