package adapters

import (
	"github.com/gofiber/fiber/v2"
	"github.com/songwaad/cs-event-backend/entities"
	"github.com/songwaad/cs-event-backend/usecases"
)

type HttpEventHandle struct {
	eventUseCase usecases.EventUseCase
}

func NewHttpEventHandle(eventUseCase usecases.EventUseCase) *HttpEventHandle {
	return &HttpEventHandle{eventUseCase: eventUseCase}
}

// Handler functions
// CreateEvent godoc
// @Summary Create Event
// @Tags Event
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 201 {array} entities.Event
// @Router /event/:id [post]
func (h *HttpEventHandle) CreateEvent(c *fiber.Ctx) error {
	var event entities.Event
	if err := c.BodyParser(&event); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	if err := h.eventUseCase.CreateEvent(&event); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(event)
}

// Handler functions
// GetEventByID godoc
// @Summary Get Event By ID
// @Tags Event
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array} entities.Event
// @Router /event/:id [get]
func (h *HttpEventHandle) GetEventByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	event, err := h.eventUseCase.GetEventByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(event)
}

// Handler functions
// GetAllEvents godoc
// @Summary Get All Events
// @Tags Event
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array} entities.Event
// @Router /events [get]
func (h *HttpEventHandle) GetAllEvents(c *fiber.Ctx) error {
	events, err := h.eventUseCase.GetAllEvents()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(events)
}

// Handler functions
// UpdateEvent godoc
// @Summary Update Event
// @Tags Event
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 201 {array} entities.Event
// @Router /event/:id [put]
func (h *HttpEventHandle) UpdateEvent(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	var event entities.Event
	if err := c.BodyParser(&event); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	event.ID = uint(id)
	if err := h.eventUseCase.UpdateEvent(&event); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Event updated successfully",
		"event":   event,
	})
}

// Handler functions
// DeleteEvent godoc
// @Summary Delete Event
// @Tags Event
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 204 {array} entities.Event
// @Router /event/:id [delete]
func (h *HttpEventHandle) DeleteEvent(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	if err := h.eventUseCase.DeleteEvent(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
