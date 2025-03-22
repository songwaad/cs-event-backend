package adapters

import (
	"github.com/gofiber/fiber/v2"
	"github.com/songwaad/cs-event-backend/entities"
	"github.com/songwaad/cs-event-backend/usecases"
)

type HttpInstructorHandle struct {
	instructorUseCase usecases.InstructorUseCase
}

func NewHttpInstructorHandle(instructorUseCase usecases.InstructorUseCase) *HttpInstructorHandle {
	return &HttpInstructorHandle{instructorUseCase: instructorUseCase}
}

func (h *HttpInstructorHandle) CreateInstructor(c *fiber.Ctx) error {
	var instructor entities.Instructor
	if err := c.BodyParser(&instructor); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	if err := h.instructorUseCase.CreateInstructor(&instructor); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(instructor)
}

func (h *HttpInstructorHandle) GetInstructorByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	instructor, err := h.instructorUseCase.GetInstructorByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(instructor)
}

// Handler functions
// GetAllInstructors godoc
// @Summary Get all Instructors
// @Description Get details of all Instructors
// @Tags Instructors
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array} entities.Instructor
// @Description Instructor struct with Gorm Model and soft delete
// @Router /Instructors [get]
func (h *HttpInstructorHandle) GetAllInstructors(c *fiber.Ctx) error {
	instructors, err := h.instructorUseCase.GetAllInstructors()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(instructors)
}

func (h *HttpInstructorHandle) UpdateInstructor(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	var instructor entities.Instructor
	if err := c.BodyParser(&instructor); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	instructor.ID = uint(id)
	if err := h.instructorUseCase.UpdateInstructor(&instructor); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(instructor)
}

func (h *HttpInstructorHandle) DeleteInstructor(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	if err := h.instructorUseCase.DeleteInstructor(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	c.Status(fiber.StatusNoContent)
	return c.JSON(fiber.Map{
		"message": "Instructor deleted successfully",
	})
}
