package adapters

import (
	"time"

	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/songwaad/cs-event-backend/entities"
	"github.com/songwaad/cs-event-backend/usecases"
)

type UserHandler struct {
	UserUseCase usecases.UserUseCase
}

func NewUserHandler(userUseCase usecases.UserUseCase) *UserHandler {
	return &UserHandler{UserUseCase: userUseCase}
}

// Handler functions
// Register godoc
// @Summary Register
// @Tags User
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 201 {array} entities.User
// @Router /auth/register [post]
func (h *UserHandler) Register(c *fiber.Ctx) error {
	var user entities.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if user.Email == "" || user.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "email and password are required",
		})
	}

	user.UserStatus.ID = 1

	if err := h.UserUseCase.Register(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user.Password = ""

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"user":    user,
	})
}

// Handler functions
// Login godoc
// @Summary Login
// @Tags User
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array} entities.User
// @Router /auth/login [post]
func (h *UserHandler) Login(c *fiber.Ctx, jwtSecretKey string) error {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid input",
		})
	}

	user, err := h.UserUseCase.Login(input.Email, input.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid email or password",
		})

	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["role"] = user.UserRole.Role
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    t,
		Expires:  time.Now().Add(time.Hour * 72),
		HTTPOnly: false,
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "login success",
		"token":   t,
	})
}

// Handler functions
// GetUserByID godoc
// @Summary Get User By ID
// @Tags User
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array} entities.User
// @Router /user/:id [get]
func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	userID := c.Params("id")

	user, err := h.UserUseCase.GetUserByID(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user.Password = ""
	return c.Status(fiber.StatusOK).JSON(user)
}

// Handler functions
// GetAllUsers godoc
// @Summary Get All Users
// @Tags User
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array} entities.User
// @Router /users [get]
func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.UserUseCase.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	for i := range users {
		users[i].Password = ""
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

// Handler functions
// Update User godoc
// @Summary Update User
// @Tags User
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 201 {array} entities.User
// @Router /user/:id [put]
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	userId := c.Params("id")

	var user entities.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user.ID = userId

	if err := h.UserUseCase.UpdateUser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	user.Password = ""
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User updated successfully",
		"user":    user,
	})
}

// Handler functions
// DeleteUser godoc
// @Summary Delete User
// @Tags User
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 204 {array} entities.User
// @Router /user/:id [delete]
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	userID := c.Params("id")

	if err := h.UserUseCase.DeleteUser(userID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// Handler functions
// Logout godoc
// @Summary Logout
// @Tags User
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array} entities.User
// @Router /auth/logout [post]
func (h *UserHandler) Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: false,
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "logout success",
	})
}
