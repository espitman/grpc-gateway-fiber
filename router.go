package main

import (
	"github.com/gofiber/fiber/v2"
)

type router struct {
    
	priceServiceHandler *priceServiceHandler
	
}

func newRouter(
    
    priceServiceHandler *priceServiceHandler,
    
) *router {
	return &router{
        
		priceServiceHandler: priceServiceHandler,
        
	}
}

func (r *router) serve(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
    
	r.priceServiceRouter(v1.Group("/price"))
	
}
