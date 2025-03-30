package adapters

import (
	"github.com/gofiber/fiber/v2"
	"github.com/songwaad/cs-event-backend/dto"
	"github.com/songwaad/cs-event-backend/usecases"
)

type HttpNotificationHandler struct {
	notificationUseCase usecases.NotificationUsecase
}

func NewHttpNotificationHandler(notificationUseCase usecases.NotificationUsecase) *HttpNotificationHandler {
	return &HttpNotificationHandler{notificationUseCase: notificationUseCase}
}

// GetNotifyByID godoc
// @Summary Get notifications by user ID
// @Description Retrieve all notifications for a specific user by their ID
// @Tags Notification
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "User ID"
// @Success 200 {array} dto.NotifyResponseDTO
// @Failure 404 {object} map[string]interface{}
// @Router /notification/user/{id} [get]
func (h *HttpNotificationHandler) GetNotifyByID(c *fiber.Ctx) error {
	userID := c.Params("id")

	notifications, err := h.notificationUseCase.GetNotifyByUserID(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var responses []dto.NotifyResponseDTO

	for _, notification := range notifications {
		response := dto.ToNotifyResponseDTP(&notification)
		responses = append(responses, response)
	}

	return c.Status(fiber.StatusOK).JSON(responses)
}

// InActive godoc
// @Summary Deactivate a notification
// @Description Deactivate a notification by setting its active status to false
// @Tags Notification
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Notification ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /notification/{id}/inactive [patch]
func (h *HttpNotificationHandler) InActive(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	if err := h.notificationUseCase.InActive(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Notification deactivated successfully",
	})
}

// DeleteNotify godoc
// @Summary Delete a notification
// @Description Delete a notification by its ID
// @Tags Notification
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Notification ID"
// @Success 204 {object} nil
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /notification/{id} [delete]
func (h *HttpNotificationHandler) DeleteNotify(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	if err := h.notificationUseCase.DeleteNotify(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
