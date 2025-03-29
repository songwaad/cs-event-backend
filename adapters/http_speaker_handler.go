package adapters

import (
	"github.com/gofiber/fiber/v2"
	"github.com/songwaad/cs-event-backend/dto"
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
// @Description Creates a new speaker by providing the speaker's first name, last name, and description
// @Tags Speaker
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param speaker body entities.Speaker true "Speaker details"
// @Success 201 {object} map[string]interface{} "Speaker created successfully"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /speaker [post]
func (h *HttpSpeakerHandle) CreateSpeaker(c *fiber.Ctx) error {
	var input entities.Speaker
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	speaker := entities.Speaker{
		SpeakerID:   input.SpeakerID,
		FirstName:   input.FirstName,
		Lastname:    input.Lastname,
		Description: input.Description,
	}

	if err := h.speakerUseCase.CreateSpeaker(&speaker); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	creatSpeaker, err := h.speakerUseCase.GetSpeakerByID(input.SpeakerID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to retrieve created Speaker",
		})
	}

	response := dto.ToSpeakerResponseDTO(*creatSpeaker)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Speaker created successfully",
		"speaker": response,
	})
}

// GetSpeakerByID godoc
// @Summary Get a Speaker by ID
// @Description Retrieve a speaker's details by their unique speaker ID
// @Tags Speaker
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Speaker ID"
// @Success 200 {object} dto.SpeakerDTO "Speaker details retrieved successfully"
// @Failure 400 {object} map[string]interface{} "Invalid ID format"
// @Failure 404 {object} map[string]interface{} "Speaker not found"
// @Router /speaker/{id} [get]
func (h *HttpSpeakerHandle) GetSpeakerByID(c *fiber.Ctx) error {
	speakerID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	speaker, err := h.speakerUseCase.GetSpeakerByID(uint(speakerID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := dto.ToSpeakerResponseDTO(*speaker)
	return c.Status(fiber.StatusOK).JSON(response)
}

// GetAllSpeakers godoc
// @Summary Get all Speakers
// @Description Retrieve a list of all speakers
// @Tags Speaker
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} dto.SpeakerDTO "List of all speakers"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /speakers [get]
func (h *HttpSpeakerHandle) GetAllSpeakers(c *fiber.Ctx) error {
	speakers, err := h.speakerUseCase.GetAllSpeakers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var responses []dto.SpeakerDTO

	for _, speaker := range speakers {
		response := dto.ToSpeakerResponseDTO(speaker)
		responses = append(responses, response)
	}

	return c.Status(fiber.StatusOK).JSON(responses)
}

// UpdateSpeaker godoc
// @Summary Update an existing Speaker
// @Description Update a speaker's details (first name, last name, description) by their unique speaker ID
// @Tags Speaker
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Speaker ID"
// @Param speaker body dto.SpeakerDTO true "Updated speaker details"
// @Success 200 {object} map[string]interface{} "Speaker updated successfully"
// @Failure 400 {object} map[string]interface{} "Invalid input or speaker ID"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /speaker/{id} [put]
func (h *HttpSpeakerHandle) UpdateSpeaker(c *fiber.Ctx) error {
	speakerID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid speaker_id",
		})
	}

	var input dto.SpeakerDTO
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	input.SpeakerID = uint(speakerID)
	speaker := entities.Speaker{
		SpeakerID:   input.SpeakerID,
		FirstName:   input.FirstName,
		Lastname:    input.Lastname,
		Description: input.Description,
	}

	if err := h.speakerUseCase.UpdateSpeaker(&speaker); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	updatedSpeaker, err := h.speakerUseCase.GetSpeakerByID(input.SpeakerID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to retrieve created Speaker",
		})
	}

	response := dto.ToSpeakerResponseDTO(*updatedSpeaker)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Speaker updated successfully",
		"speaker": response,
	})
}

// DeleteSpeaker godoc
// @Summary Delete a speaker
// @Description Delete a speaker by their unique speaker ID
// @Tags Speaker
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Speaker ID"
// @Success 204 {object} nil "Speaker deleted successfully"
// @Failure 400 {object} map[string]interface{} "Invalid speaker ID"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /speaker/{id} [delete]
func (h *HttpSpeakerHandle) DeleteSpeaker(c *fiber.Ctx) error {
	speakerID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	if err := h.speakerUseCase.DeleteSpeaker(uint(speakerID)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
