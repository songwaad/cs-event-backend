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

func (h *HttpEventHandle) GetAllEvents(c *fiber.Ctx) error {
	events, err := h.eventUseCase.GetAllEvents()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(events)
}

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

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"message": "Event deleted successfully",
	})
}
