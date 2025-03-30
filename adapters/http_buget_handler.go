package adapters

import (
	"github.com/gofiber/fiber/v2"
	"github.com/songwaad/cs-event-backend/dto"
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
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Budget body dto.BudgetDTO true "Budget object"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /budget [post]
func (h *HttpBudgetHandle) CreateBudget(c *fiber.Ctx) error {
	var input dto.BudgetDTO
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	budget := entities.Budget{
		EventID:     input.EventID,
		Amount:      input.Amount,
		Description: input.Description,
	}

	if err := h.budgetUseCase.CreateBudget(&budget); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	createBudget, err := h.budgetUseCase.GetBudgetByID(budget.BudgetID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to retrieve created Budget",
		})
	}

	response := dto.ToBudgetResponseDTO(createBudget)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Budget created successfully",
		"budget":  response,
	})
}

// GetBudgetByID godoc
// @Summary Get a Budget by ID
// @Description Retrieve a Budget by ID
// @Tags Budget
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Budget ID"
// @Success 200 {object} dto.BudgetDTO
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

	budget, err := h.budgetUseCase.GetBudgetByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := dto.ToBudgetResponseDTO(budget)
	return c.Status(fiber.StatusOK).JSON(response)
}

// GetAllBudgets godoc
// @Summary Get all Budgets
// @Description Retrieve a list of all Budgets
// @Tags Budget
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} dto.BudgetDTO
// @Failure 500 {object} map[string]interface{}
// @Router /budgets [get]
func (h *HttpBudgetHandle) GetAllBudgets(c *fiber.Ctx) error {
	budgets, err := h.budgetUseCase.GetALLBudgets()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var responses []dto.BudgetDTO

	for _, budget := range budgets {
		response := dto.ToBudgetResponseDTO(&budget)
		responses = append(responses, response)
	}

	return c.Status(fiber.StatusOK).JSON(responses)
}

// UpdateBudget godoc
// @Summary Update an existing Budget
// @Description Update an existing Budget by ID
// @Tags Budget
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Budget ID"
// @Param Budget body dto.BudgetDTO true "Updated Budget object"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /budgets/{id} [put]
func (h *HttpBudgetHandle) UpdateBudget(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	var input dto.BudgetDTO
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	input.BudgetID = uint(id)
	budget := entities.Budget{
		BudgetID:    input.BudgetID,
		EventID:     input.EventID,
		Amount:      input.Amount,
		Description: input.Description,
	}
	if err := h.budgetUseCase.UpdateBudget(&budget); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	updatedBudget, err := h.budgetUseCase.GetBudgetByID(input.BudgetID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to retrieve updated Speaker",
		})
	}

	response := dto.ToBudgetResponseDTO(updatedBudget)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "budget updated successfully",
		"budget":  response,
	})
}

// DeleteBudget godoc
// @Summary Delete a Budget
// @Description Delete a Budget by ID
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
