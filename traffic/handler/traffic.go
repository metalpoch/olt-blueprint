package handler

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/metalpoch/olt-blueprint/common/model"
	"github.com/metalpoch/olt-blueprint/common/usecase"
)

type TrafficHandler struct {
	Usecase usecase.TrafficUsecase
}

// Traffic by interface
//
//	@Summary		Get a array of interface traffic in a range date
//	@Description	get traffic by interface ID
//	@Tags			traffic
//	@Produce		json
//	@Param			id	path		uint	true	"Interface ID"
//	@Success		200	{object}	[]model.Traffic
//	@Failure		400	{object}	object{message=string}
//	@Failure		500	{object}	object{message=string}
//	@Router			/interface/{id} [get]
func (hdlr TrafficHandler) GetByInterface(c fiber.Ctx) error {
	id, err := fiber.Convert(c.Params("id"), strconv.Atoi)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	query := new(model.TranficRangeDate)
	if err := c.Bind().Query(query); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	res, err := hdlr.Usecase.GetTrafficByDevice(uint(id), query)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(res)
}

// Traffic by device
//
//	@Summary		Get a array of device traffic in a range date
//	@Description	get traffic by device ID
//	@Tags			traffic
//	@Produce		json
//	@Param			id	path		uint	true	"Device ID"
//	@Success		200	{object}	[]model.Traffic
//	@Failure		400	{object}	object{message=string}
//	@Failure		500	{object}	object{message=string}
//	@Router			/device/{id} [get]
func (hdlr TrafficHandler) GetByDevice(c fiber.Ctx) error {
	id, err := fiber.Convert(c.Params("id"), strconv.Atoi)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	query := new(model.TranficRangeDate)
	if err := c.Bind().Query(query); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	res, err := hdlr.Usecase.GetTrafficByDevice(uint(id), query)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(res)
}