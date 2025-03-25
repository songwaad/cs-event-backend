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

// CreateEvent godoc
// @Summary Create a new event
// @Description Create a new event with the provided details
// @Tags Event
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param event body entities.Event true "Event object"
// @Success 201 {object} entities.Event
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /event [post]
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

// GetEventByID godoc
// @Summary Get an event by ID
// @Description Retrieve an event by its ID
// @Tags Event
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Event ID"
// @Success 200 {object} entities.Event
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /event/{id} [get]
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

// GetAllEvents godoc
// @Summary Get all events
// @Description Retrieve a list of all events
// @Tags Event
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} entities.Event
// @Failure 500 {object} map[string]interface{}
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

// UpdateEvent godoc
// @Summary Update an event
// @Description Update an existing event by ID
// @Tags Event
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Event ID"
// @Param event body entities.Event true "Updated event object"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /event/{id} [put]
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

// DeleteEvent godoc
// @Summary Delete an event
// @Description Delete an event by ID
// @Tags Event
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Event ID"
// @Success 204 {object} nil
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /event/{id} [delete]
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
