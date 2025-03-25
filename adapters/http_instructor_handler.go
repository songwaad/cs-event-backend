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

// CreateInstructor godoc
// @Summary Create a new instructor
// @Description Create a new instructor with the provided details
// @Tags Instructor
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param instructor body entities.Instructor true "Instructor object"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /instructor [post]
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

// GetInstructorByID godoc
// @Summary Get an instructor by ID
// @Description Retrieve an instructor by its ID
// @Tags Instructor
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Instructor ID"
// @Success 200 {object} entities.Instructor
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /instructor/{id} [get]
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

	return c.Status(fiber.StatusOK).JSON(instructor)
}

// GetAllInstructors godoc
// @Summary Get all instructors
// @Description Retrieve a list of all instructors
// @Tags Instructor
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} entities.Instructor
// @Failure 500 {object} map[string]interface{}
// @Router /instructors [get]
func (h *HttpInstructorHandle) GetAllInstructors(c *fiber.Ctx) error {
	instructors, err := h.instructorUseCase.GetAllInstructors()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(instructors)
}

// UpdateInstructor godoc
// @Summary Update an instructor
// @Description Update an existing instructor by ID
// @Tags Instructor
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Instructor ID"
// @Param instructor body entities.Instructor true "Updated instructor object"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /instructor/{id} [put]
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

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "Instructor updated successfully",
		"instructor": instructor,
	})
}

// DeleteInstructor godoc
// @Summary Delete an instructor
// @Description Delete an instructor by ID
// @Tags Instructor
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Instructor ID"
// @Success 204 {object} nil
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /instructor/{id} [delete]
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

	return c.SendStatus(fiber.StatusNoContent)
}
