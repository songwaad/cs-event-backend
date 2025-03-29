package adapters

import (
	"github.com/gofiber/fiber/v2"
	"github.com/songwaad/cs-event-backend/dto"
	"github.com/songwaad/cs-event-backend/usecases"
)

type HttpStrategyHandler struct {
	strategyUsecase usecases.StrategyUsecases
}

func NewHttpStrategyHandler(strategyUseCase usecases.StrategyUsecases) *HttpStrategyHandler {
	return &HttpStrategyHandler{strategyUsecase: strategyUseCase}
}

// GetStrategyByID godoc
// @Summary Get a strategy by ID
// @Description Retrieve a strategy by its unique ID from the database
// @Tags Strategy
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Strategy ID"
// @Success 200 {object} dto.StrategyResponseDTO "Successfully retrieved strategy"
// @Failure 400 {object} map[string]interface{} "Invalid ID format"
// @Failure 404 {object} map[string]interface{} "Strategy not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /strategy/{id} [get]
func (h *HttpStrategyHandler) GetStrategyByID(c *fiber.Ctx) error {
	strategyID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	strategy, err := h.strategyUsecase.GetStrategyByID(uint(strategyID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := dto.ToStrategyResponseDTO(*strategy)
	return c.Status(fiber.StatusOK).JSON(response)
}

// GetAllStrategies godoc
// @Summary Get all strategies
// @Description Retrieve a list of all strategies
// @Tags Strategy
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} dto.StrategyResponseDTO "List of all strategies"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /strategies [get]
func (h *HttpStrategyHandler) GetAllStrategies(c *fiber.Ctx) error {
	strategies, err := h.strategyUsecase.GetAllStrategy()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var responses []dto.StrategyResponseDTO
	for _, strategy := range strategies {
		response := dto.ToStrategyResponseDTO(strategy)
		responses = append(responses, response)
	}

	return c.Status(fiber.StatusOK).JSON(responses)
}

// GetEventStrategyByID godoc
// @Summary Get an event strategy by ID
// @Description Retrieve an event strategy by its unique ID
// @Tags Event Strategy
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Event Strategy ID"
// @Success 200 {object} dto.EventStrategyResponseDTO "Successfully retrieved event strategy"
// @Failure 400 {object} map[string]interface{} "Invalid ID format"
// @Failure 404 {object} map[string]interface{} "Event strategy not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /event-strategy/{id} [get]
func (h *HttpStrategyHandler) GetEventStrategyByID(c *fiber.Ctx) error {
	eventStrategyID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	eventStrategy, err := h.strategyUsecase.GetEventeStrategyByID(uint(eventStrategyID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := dto.ToEventStrategyResponseDTO(*eventStrategy)
	return c.Status(fiber.StatusOK).JSON(response)
}

// GetAllEventStrategies godoc
// @Summary Get all event strategies
// @Description Retrieve a list of all event strategies
// @Tags Event Strategy
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} dto.EventStrategyResponseDTO "List of all event strategies"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /event-strategies [get]
func (h *HttpStrategyHandler) GetAllEventStrategies(c *fiber.Ctx) error {
	eventStrategies, err := h.strategyUsecase.GetAllEventeStrategy()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var responses []dto.EventStrategyResponseDTO
	for _, eventStrategy := range eventStrategies {
		response := dto.ToEventStrategyResponseDTO(eventStrategy)
		responses = append(responses, response)
	}

	return c.Status(fiber.StatusOK).JSON(responses)
}
