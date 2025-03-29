package adapters

import (
	"time"

	"github.com/gofiber/fiber/v2"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/songwaad/cs-event-backend/dto"
	"github.com/songwaad/cs-event-backend/entities"
	"github.com/songwaad/cs-event-backend/usecases"
)

type UserHandler struct {
	UserUseCase usecases.UserUseCase
}

func NewUserHandler(userUseCase usecases.UserUseCase) *UserHandler {
	return &UserHandler{UserUseCase: userUseCase}
}

// Register godoc
// @Summary Register a new user
// @Description Registers a new user by providing an email, password, and optional first and last names, as well as user role ID.
// @Tags User
// @Accept json
// @Produce json
// @Param user body dto.UserRegisterDTO true "User registration details"
// @Success 201 {object} dto.UserRegisterDTO "User successfully created"
// @Failure 400 {object} map[string]interface{} "Bad request due to invalid or missing input"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /auth/register [post]
func (h *UserHandler) Register(c *fiber.Ctx) error {
	var input dto.UserRegisterDTO
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if input.Email == "" || input.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "email and password are required",
		})
	}

	user := entities.User{
		UserID:       input.UserID,
		Email:        input.Email,
		Password:     input.Password,
		FirstName:    input.FirstName,
		Lastname:     input.Lastname,
		UserRoleID:   input.UserRoleID,
		UserStatusID: 1,
	}

	if err := h.UserUseCase.Register(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	createdUser, err := h.UserUseCase.GetUserByID(user.UserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to retrieve created user",
		})
	}

	response := dto.ToUserResponseDTO(*createdUser)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"user":    response,
	})
}

// Login godoc
// @Summary User login
// @Description Authenticates a user by their email and password, and returns a JWT token upon successful authentication.
// @Tags User
// @Accept json
// @Produce json
// @Param credentials body object{email=string,password=string} true "Login credentials"
// @Success 200 {object} map[string]interface{} "Login successful with JWT token"
// @Failure 400 {object} map[string]interface{} "Bad request due to invalid credentials"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /auth/login [post]
func (h *UserHandler) Login(c *fiber.Ctx, jwtSecretKey string) error {
	var input dto.UserLoginDTO
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
	claims["user_id"] = user.UserRoleID
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

// GetUserByID godoc
// @Summary Retrieve a user by ID
// @Description Fetches the user details based on their unique user ID.
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "User ID"
// @Success 200 {object} dto.UserResponseDTO "User details retrieved successfully"
// @Failure 404 {object} map[string]interface{} "User not found"
// @Router /user/{id} [get]
func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	userID := c.Params("id")

	user, err := h.UserUseCase.GetUserByID(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	response := dto.ToUserResponseDTO(*user)
	return c.Status(fiber.StatusOK).JSON(response)
}

// GetAllUsers godoc
// @Summary Retrieve a list of all users
// @Description Fetches a list of all registered users with their details.
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} dto.UserResponseDTO "List of users retrieved successfully"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /users [get]
func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.UserUseCase.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var responses []dto.UserResponseDTO

	for _, user := range users {
		response := dto.ToUserResponseDTO(user)
		responses = append(responses, response)
	}

	return c.Status(fiber.StatusOK).JSON(responses)
}

// ChangePassword godoc
// @Summary Change user password
// @Description Allows a user to change their password after verifying the old password.
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "User ID"
// @Param password body dto.UserChangePasswordDTO true "Old and new password"
// @Success 200 {object} map[string]interface{} "Password successfully changed"
// @Failure 400 {object} map[string]interface{} "Bad request due to invalid input"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /user/{id}/change-password [put]
func (h *UserHandler) ChangePassword(c *fiber.Ctx) error {
	userID := c.Params("id")

	var input dto.UserChangePasswordDTO
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid input: " + err.Error(),
		})
	}

	if input.OldPassword == "" || input.NewPassword == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "old and new password are required",
		})
	}

	if err := h.UserUseCase.ChangePassword(userID, input.OldPassword, input.NewPassword); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Password changed successfully",
	})
}

// UpdateRole godoc
// @Summary Update user role
// @Description Updates the role of a user by their unique ID.
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "User ID"
// @Param role body dto.UserUpdateUserRoleDTO true "New user role"
// @Success 200 {object} dto.UserResponseDTO "User role updated successfully"
// @Failure 400 {object} map[string]interface{} "Bad request due to invalid input"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /user/{id}/role [put]
func (h *UserHandler) UpdateRole(c *fiber.Ctx) error {
	userID := c.Params("id")
	var input dto.UserUpdateUserRoleDTO
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid input: " + err.Error(),
		})
	}

	if input.UserRoleID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "user_role_id is required",
		})
	}

	if err := h.UserUseCase.UpdateUserRole(userID, uint(input.UserRoleID)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to update role: " + err.Error(),
		})
	}

	user, _ := h.UserUseCase.GetUserByID(userID) // ดึงข้อมูลล่าสุด
	response := dto.ToUserResponseDTO(*user)
	return c.Status(fiber.StatusOK).JSON(response)
}

// UpdateStatus godoc
// @Summary Update user status
// @Description Changes the user status (e.g., active, inactive) based on their ID.
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "User ID"
// @Param status body dto.UserUpdateUserStatusDTO true "New user status"
// @Success 200 {object} dto.UserResponseDTO "User status updated successfully"
// @Failure 400 {object} map[string]interface{} "Bad request due to invalid input"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /user/{id}/status [put]
func (h *UserHandler) UpdateStatus(c *fiber.Ctx) error {
	userID := c.Params("id")
	var input dto.UserUpdateUserStatusDTO
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid input: " + err.Error(),
		})
	}

	if input.UserStatusID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "user_status_id is required",
		})
	}

	if err := h.UserUseCase.UpdateUserStatus(userID, input.UserStatusID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to update status: " + err.Error(),
		})
	}

	user, _ := h.UserUseCase.GetUserByID(userID) // ดึงข้อมูลล่าสุด
	response := dto.ToUserResponseDTO(*user)
	return c.Status(fiber.StatusOK).JSON(response)
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Deletes a user from the system using their unique ID.
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "User ID"
// @Success 204 {object} nil "User successfully deleted"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /user/{id} [delete]
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	userID := c.Params("id")

	if err := h.UserUseCase.DeleteUser(userID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// Logout godoc
// @Summary User logout
// @Description Logs out the user by clearing their JWT cookie.
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{} "User logged out successfully"
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
