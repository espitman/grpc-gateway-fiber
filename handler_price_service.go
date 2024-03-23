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

// V1CalendarGet
// @Summary V1CalendarGet
// @Description V1CalendarGet
// @Tags price_service
// @Produce json
// @Param startDate query string false "startDate"
// @Param endDate query string false "endDate"
// @Param limit query int false "limit"
// @Param offset query int false "offset"
// @Success 200 {object} pricepb.V1CalendarGetResponse
// @Router /api/v1/price/calendar [Get]
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

// V1DayDisable
// @Summary V1DayDisable
// @Description V1DayDisable
// @Tags price_service
// @Produce json
// @Param body body pricepb.V1DayEnableDisableRequest true "body"
// @Success 200 {object} pricepb.V1DayEnableDisableResponse
// @Router /api/v1/price/V1DayDisable [Post]
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

// V1DayEnable
// @Summary V1DayEnable
// @Description V1DayEnable
// @Tags price_service
// @Produce json
// @Param body body pricepb.V1DayEnableDisableRequest true "body"
// @Success 200 {object} pricepb.V1DayEnableDisableResponse
// @Router /api/v1/price/V1DayEnable [Post]
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

// V1DayGuarantee
// @Summary V1DayGuarantee
// @Description V1DayGuarantee
// @Tags price_service
// @Produce json
// @Param body body pricepb.V1DayGuaranteeRequest true "body"
// @Success 200 {object} pricepb.V1DayGuaranteeResponse
// @Router /api/v1/price/V1DayGuarantee [Post]
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

// V1DayUnGuarantee
// @Summary V1DayUnGuarantee
// @Description V1DayUnGuarantee
// @Tags price_service
// @Produce json
// @Param body body pricepb.V1DayUnGuaranteeRequest true "body"
// @Success 200 {object} pricepb.V1DayUnGuaranteeResponse
// @Router /api/v1/price/V1DayUnGuarantee [Post]
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

// V1PriceCreate
// @Summary V1PriceCreate
// @Description V1PriceCreate
// @Tags price_service
// @Produce json
// @Param body body pricepb.V1PriceCreateRequest true "body"
// @Success 200 {object} pricepb.V1PriceGetResponse
// @Router /api/v1/price/V1PriceCreate [Post]
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

// V1PriceGet
// @Summary V1PriceGet
// @Description V1PriceGet
// @Tags price_service
// @Produce json
// @Param accommodationID path string false "accommodationID"
// @Success 200 {object} pricepb.V1PriceGetResponse
// @Router /api/v1/price/price/{accommodationID} [Get]
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

// V1PriceUpdate
// @Summary V1PriceUpdate
// @Description V1PriceUpdate
// @Tags price_service
// @Produce json
// @Param body body pricepb.V1PriceCreateRequest true "body"
// @Success 200 {object} pricepb.V1PriceGetResponse
// @Router /api/v1/price/V1PriceUpdate [Post]
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
