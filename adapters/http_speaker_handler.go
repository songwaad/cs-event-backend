package adapters

import (
	"github.com/gofiber/fiber/v2"
	"github.com/songwaad/cs-event-backend/entities"
	"github.com/songwaad/cs-event-backend/usecases"
)

type HttpSpeakerHandle struct {
	speakerUseCase usecases.SpeakerUseCase
}

func NewHttpSpeakerHandle(speakerUseCase usecases.SpeakerUseCase) *HttpSpeakerHandle {
	return &HttpSpeakerHandle{speakerUseCase: speakerUseCase}
}

// CreateSpeaker godoc
// @Summary Create a new Speaker
// @Description Create a new Speaker with the provided details
// @Tags Speaker
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param Speaker body entities.Speaker true "Speaker object"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /speaker [post]
func (h *HttpSpeakerHandle) CreateSpeaker(c *fiber.Ctx) error {
	var instructor entities.Speaker
	if err := c.BodyParser(&instructor); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	if err := h.speakerUseCase.CreateSpeaker(&instructor); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(instructor)
}

// GetSpeakerByID godoc
// @Summary Get an Speaker by ID
// @Description Retrieve an Speaker by its ID
// @Tags Speaker
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Speaker ID"
// @Success 200 {object} entities.Speaker
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /speaker/{id} [get]
func (h *HttpSpeakerHandle) GetSpeakerByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	speaker, err := h.speakerUseCase.GetSpeakerByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(speaker)
}

// GetAllSpeakers godoc
// @Summary Get all Speakers
// @Description Retrieve a list of all Speakers
// @Tags Speakers
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} entities.Speaker
// @Failure 500 {object} map[string]interface{}
// @Router /speakers [get]
func (h *HttpSpeakerHandle) GetAllSpeakers(c *fiber.Ctx) error {
	speakers, err := h.speakerUseCase.GetAllSpeakers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(speakers)
}

// UpdateSpeaker godoc
// @Summary Update an Speaker
// @Description Update an existing Speaker by ID
// @Tags Speaker
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Speaker ID"
// @Param Speaker body entities.Speaker true "Updated Speaker object"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /speaker/{id} [put]
func (h *HttpSpeakerHandle) UpdateSpeaker(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	var speaker entities.Speaker
	if err := c.BodyParser(&speaker); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	speaker.ID = uint(id)
	if err := h.speakerUseCase.UpdateSpeaker(&speaker); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Speaker updated successfully",
		"speaker": speaker,
	})
}

// DeleteSpeaker godoc
// @Summary Delete an speaker
// @Description Delete an speaker by ID
// @Tags Speaker
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Speaker ID"
// @Success 204 {object} nil
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /speaker/{id} [delete]
func (h *HttpSpeakerHandle) DeleteSpeaker(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	if err := h.speakerUseCase.DeleteSpeaker(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
