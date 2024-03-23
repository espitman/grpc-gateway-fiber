package main

import (
	pricepb "git.alibaba.ir/taraaz/salvation2/monorepo/pkg/protos/protogen/price_service"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
)

type priceServiceHandler struct {
	pb pricepb.PriceServiceClient
}

func newPriceServiceHandler(pb pricepb.PriceServiceClient) *priceServiceHandler {
	return &priceServiceHandler{
		pb: pb,
	}
}

func (h *priceServiceHandler) V1CalendarGet(c *fiber.Ctx) error {
	header := metadata.New(map[string]string{"authorization": "saeed"})
	ctx := metadata.NewOutgoingContext(c.Context(), header)
	var reqDto pricepb.V1CalendarGetRequest
	_ = c.QueryParser(&reqDto)

	res, err := h.pb.V1CalendarGet(ctx, &reqDto)
	if err != nil {
		st, _ := status.FromError(err)
		log.Println("err:", st.Code(), st.Message(), st.Details())
		return c.Status(fiber.StatusBadRequest).Send(nil)
	}
	return c.JSON(res)
}

func (h *priceServiceHandler) V1DayDisable(c *fiber.Ctx) error {
	header := metadata.New(map[string]string{"authorization": "saeed"})
	ctx := metadata.NewOutgoingContext(c.Context(), header)
	var reqDto pricepb.V1DayEnableDisableRequest
	_ = c.QueryParser(&reqDto)
	_ = c.ParamsParser(&reqDto)
	_ = c.BodyParser(&reqDto)
	res, err := h.pb.V1DayDisable(ctx, &reqDto)
	if err != nil {
		st, _ := status.FromError(err)
		log.Println("err:", st.Code(), st.Message(), st.Details())
		return c.Status(fiber.StatusBadRequest).Send(nil)
	}
	return c.JSON(res)
}

func (h *priceServiceHandler) V1DayEnable(c *fiber.Ctx) error {
	header := metadata.New(map[string]string{"authorization": "saeed"})
	ctx := metadata.NewOutgoingContext(c.Context(), header)
	var reqDto pricepb.V1DayEnableDisableRequest
	_ = c.QueryParser(&reqDto)
	_ = c.ParamsParser(&reqDto)
	_ = c.BodyParser(&reqDto)
	res, err := h.pb.V1DayEnable(ctx, &reqDto)
	if err != nil {
		st, _ := status.FromError(err)
		log.Println("err:", st.Code(), st.Message(), st.Details())
		return c.Status(fiber.StatusBadRequest).Send(nil)
	}
	return c.JSON(res)
}

func (h *priceServiceHandler) V1DayGuarantee(c *fiber.Ctx) error {
	header := metadata.New(map[string]string{"authorization": "saeed"})
	ctx := metadata.NewOutgoingContext(c.Context(), header)
	var reqDto pricepb.V1DayGuaranteeRequest
	_ = c.QueryParser(&reqDto)
	_ = c.ParamsParser(&reqDto)
	_ = c.BodyParser(&reqDto)
	res, err := h.pb.V1DayGuarantee(ctx, &reqDto)
	if err != nil {
		st, _ := status.FromError(err)
		log.Println("err:", st.Code(), st.Message(), st.Details())
		return c.Status(fiber.StatusBadRequest).Send(nil)
	}
	return c.JSON(res)
}

func (h *priceServiceHandler) V1DayUnGuarantee(c *fiber.Ctx) error {
	header := metadata.New(map[string]string{"authorization": "saeed"})
	ctx := metadata.NewOutgoingContext(c.Context(), header)
	var reqDto pricepb.V1DayUnGuaranteeRequest
	_ = c.QueryParser(&reqDto)
	_ = c.ParamsParser(&reqDto)
	_ = c.BodyParser(&reqDto)
	res, err := h.pb.V1DayUnGuarantee(ctx, &reqDto)
	if err != nil {
		st, _ := status.FromError(err)
		log.Println("err:", st.Code(), st.Message(), st.Details())
		return c.Status(fiber.StatusBadRequest).Send(nil)
	}
	return c.JSON(res)
}

func (h *priceServiceHandler) V1PriceCreate(c *fiber.Ctx) error {
	header := metadata.New(map[string]string{"authorization": "saeed"})
	ctx := metadata.NewOutgoingContext(c.Context(), header)
	var reqDto pricepb.V1PriceCreateRequest
	_ = c.QueryParser(&reqDto)
	_ = c.ParamsParser(&reqDto)
	_ = c.BodyParser(&reqDto)
	res, err := h.pb.V1PriceCreate(ctx, &reqDto)
	if err != nil {
		st, _ := status.FromError(err)
		log.Println("err:", st.Code(), st.Message(), st.Details())
		return c.Status(fiber.StatusBadRequest).Send(nil)
	}
	return c.JSON(res)
}

func (h *priceServiceHandler) V1PriceGet(c *fiber.Ctx) error {
	header := metadata.New(map[string]string{"authorization": "saeed"})
	ctx := metadata.NewOutgoingContext(c.Context(), header)
	var reqDto pricepb.V1PriceGetRequest

	_ = c.ParamsParser(&reqDto)

	res, err := h.pb.V1PriceGet(ctx, &reqDto)
	if err != nil {
		st, _ := status.FromError(err)
		log.Println("err:", st.Code(), st.Message(), st.Details())
		return c.Status(fiber.StatusBadRequest).Send(nil)
	}
	return c.JSON(res)
}

func (h *priceServiceHandler) V1PriceUpdate(c *fiber.Ctx) error {
	header := metadata.New(map[string]string{"authorization": "saeed"})
	ctx := metadata.NewOutgoingContext(c.Context(), header)
	var reqDto pricepb.V1PriceCreateRequest
	_ = c.QueryParser(&reqDto)
	_ = c.ParamsParser(&reqDto)
	_ = c.BodyParser(&reqDto)
	res, err := h.pb.V1PriceUpdate(ctx, &reqDto)
	if err != nil {
		st, _ := status.FromError(err)
		log.Println("err:", st.Code(), st.Message(), st.Details())
		return c.Status(fiber.StatusBadRequest).Send(nil)
	}
	return c.JSON(res)
}
