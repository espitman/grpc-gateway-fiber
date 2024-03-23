package main

import "github.com/gofiber/fiber/v2"

func (r *router) priceServiceRouter(v fiber.Router) {

	v.Get("/calendar", r.priceServiceHandler.V1CalendarGet)

	v.Post("/V1DayDisable", r.priceServiceHandler.V1DayDisable)

	v.Post("/V1DayEnable", r.priceServiceHandler.V1DayEnable)

	v.Post("/V1DayGuarantee", r.priceServiceHandler.V1DayGuarantee)

	v.Post("/V1DayUnGuarantee", r.priceServiceHandler.V1DayUnGuarantee)

	v.Post("/V1PriceCreate", r.priceServiceHandler.V1PriceCreate)

	v.Get("/price/:accommodationID", r.priceServiceHandler.V1PriceGet)

	v.Post("/V1PriceUpdate", r.priceServiceHandler.V1PriceUpdate)

}
