package adapters

import (
	"github.com/gofiber/fiber/v2"
	"github.com/songwaad/cs-event-backend/entities"
	"github.com/songwaad/cs-event-backend/usecases"
)

type HttpBudgetHandle struct {
	budgetUseCase usecases.BudgetUseCase
}

func NewHttpBudgetHandler(budgetUseCase usecases.BudgetUseCase) *HttpBudgetHandle {
	return &HttpBudgetHandle{budgetUseCase: budgetUseCase}
}

// CreateBudget godoc
// @Summary Create a new Budget
// @Description Create a new Budget with the provided details
// @Tags Budget
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param Budget body entities.Budget true "Budget object"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /budget [post]
func (h *HttpBudgetHandle) CreateBudget(c *fiber.Ctx) error {
	var budget entities.Budget
	if err := c.BodyParser(&budget); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	if err := h.budgetUseCase.CreateBudget(&budget); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(budget)
}

// GetBudgetByID godoc
// @Summary Get an Budget by ID
// @Description Retrieve an Budget by ID
// @Tags Budget
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Budget ID"
// @Success 200 {object} entities.Budget
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /budget/{id} [get]
func (h *HttpBudgetHandle) GetBudgetByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	instructor, err := h.budgetUseCase.GetBudgetByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(instructor)
}

// GetAllBudgets godoc
// @Summary Get all Budget
// @Description Retrieve a list of all Budget
// @Tags Budget
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} entities.Budget
// @Failure 500 {object} map[string]interface{}
// @Router /budgets [get]
func (h *HttpBudgetHandle) GetAllBudgets(c *fiber.Ctx) error {
	instructors, err := h.budgetUseCase.GetALLBudgets()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(instructors)
}

// UpdateBudget godoc
// @Summary Update an instructor
// @Description Update an existing Budget by ID
// @Tags Budget
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Budget ID"
// @Param Budget body entities.Budget true "Updated instructor object"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /ิudgets/{id} [put]
func (h *HttpBudgetHandle) UpdateBudget(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	var budget entities.Budget
	if err := c.BodyParser(&budget); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	budget.ID = uint(id)
	if err := h.budgetUseCase.UpdateBudget(&budget); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "Instructor updated successfully",
		"instructor": budget,
	})
}

// BudgetInstructor godoc
// @Summary Delete an Budget
// @Description Delete an Budget by ID
// @Tags Budget
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Budget ID"
// @Success 204 {object} nil
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /budget/{id} [delete]
func (h *HttpBudgetHandle) DeleteBudget(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	if err := h.budgetUseCase.DeleteBudget(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
