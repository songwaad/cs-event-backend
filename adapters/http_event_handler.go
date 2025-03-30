package adapters

import (
	"sort"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/songwaad/cs-event-backend/dto"
	"github.com/songwaad/cs-event-backend/entities"
	"github.com/songwaad/cs-event-backend/usecases"
)

type HttpEventHandler struct {
	eventUseCase  usecases.EventUseCase
	userUseCase   usecases.UserUseCase
	notifyUseCase usecases.NotificationUsecase
}

func NewHttpEventHandler(eventUseCase usecases.EventUseCase, userUseCase usecases.UserUseCase, notifyUseCase usecases.NotificationUsecase) *HttpEventHandler {
	return &HttpEventHandler{
		eventUseCase:  eventUseCase,
		userUseCase:   userUseCase,
		notifyUseCase: notifyUseCase,
	}
}

// GetAllEventTypeStatus godoc
// @Summary Get all event type status
// @Description Retrieve all event type status from the system
// @Tags Event
// @Accept  json
// @Produce  json
// @Success 200 {array} entities.EventTypeStatus "Successfully retrieved all event type status"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /events/types/status [get]
func (h *HttpEventHandler) GetAllEventTypeStatus(c *fiber.Ctx) error {
	eventTypeStatus, err := h.eventUseCase.GetAllEventTypeStatus()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(eventTypeStatus)
}

// GetAllEventType godoc
// @Summary Get all event types
// @Description Retrieve all event types from the system
// @Tags Event
// @Accept  json
// @Produce  json
// @Success 200 {array} entities.EventType "Successfully retrieved all event types"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /events/types [get]
func (h *HttpEventHandler) GetAllEventType(c *fiber.Ctx) error {
	eventTypes, err := h.eventUseCase.GetAllEventType()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(eventTypes)
}

// GetAllEventStatus godoc
// @Summary Get all event status
// @Description Retrieve all event status from the system
// @Tags Event
// @Accept  json
// @Produce  json
// @Success 200 {array} entities.EventStatus "Successfully retrieved all event status"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /events/status [get]
func (h *HttpEventHandler) GetAllEventStatus(c *fiber.Ctx) error {
	eventStatus, err := h.eventUseCase.GetAllEventStatus()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(eventStatus)
}

// GetAllEventPlan godoc
// @Summary Get all event plans
// @Description Retrieve all event plans from the system
// @Tags Event
// @Accept  json
// @Produce  json
// @Success 200 {array} entities.EventPlan "Successfully retrieved all event plans"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /events/plans [get]
func (h *HttpEventHandler) GetAllEventPlan(c *fiber.Ctx) error {
	eventPlans, err := h.eventUseCase.GetAllEventPlan()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(eventPlans)
}

// CreateEvent godoc
// @Summary Create a new event
// @Description Create a new event in the system
// @Tags Event
// @Accept json
// @Produce json
// @Param event body dto.EventCreateDTO true "Event information"
// @Success 201 {object} dto.EventResponseDTO "Successfully created event"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /event [post]
func (h *HttpEventHandler) CreateEvent(c *fiber.Ctx) error {
	var input dto.EventCreateDTO
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	event := entities.Event{
		Name:              input.Name,
		Year:              input.Year,
		Rationale:         input.Rationale,
		Objective:         input.Objective,
		StartDate:         input.StartDate,
		EndDate:           input.EndDate,
		Location:          input.Location,
		Methodology:       input.Methodology,
		HasBudget:         input.HasBudget,
		Monitoring:        input.Monitoring,
		EventTypeStatusID: input.EventTypeStatusID,
		EventPlanID:       input.EventPlanID,
		EventTypeID:       input.EventTypeID,
		EventStrategyID:   input.EventStrategyID,
		CreatedByUserID:   input.CreatedByUserID,
		EventStatusID:     input.EventStatusID,
	}

	// สร้าง event ใหม่
	if err := h.eventUseCase.CreateEvent(&event); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// เชื่อมโยง Responsible Users
	if len(input.ResponsibleUserIDs) > 0 {
		if err := h.eventUseCase.AddResponsibleUsersToEvent(&event, input.ResponsibleUserIDs); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
	}

	// เชื่อมโยง Speakers
	if len(input.SpeakerIDs) > 0 {
		if err := h.eventUseCase.AddSpeakersToEvent(&event, input.SpeakerIDs); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
	}

	createdEvent, err := h.eventUseCase.GetEventByID(uint(event.EventID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to retrieve created event",
		})
	}

	response := dto.ToEventResponseDTO(createdEvent)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Event created successfully",
		"event":   response,
	})
}

// GetEventByID godoc
// @Summary Get event by ID
// @Description Retrieve an event by its ID
// @Tags Event
// @Accept json
// @Produce json
// @Param id path int true "Event ID"
// @Success 200 {object} dto.EventResponseDTO "Successfully retrieved event"
// @Failure 404 {object} map[string]interface{} "Event not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /event/{id} [get]
func (h *HttpEventHandler) GetEventByID(c *fiber.Ctx) error {
	eventID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid event ID",
		})
	}

	event, err := h.eventUseCase.GetEventByID(uint(eventID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Event not found",
		})
	}

	response := dto.ToEventResponseDTO(event)
	return c.Status(fiber.StatusOK).JSON(response)
}

// GetAllEvents godoc
// @Summary Get all events
// @Description Retrieve all events from the system
// @Tags Event
// @Accept json
// @Produce json
// @Success 200 {array} dto.EventResponseDTO "Successfully retrieved all events"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /events [get]
func (h *HttpEventHandler) GetAllEvents(c *fiber.Ctx) error {
	events, err := h.eventUseCase.GetAllEvents()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var responses []dto.EventResponseDTO

	for _, event := range events {
		response := dto.ToEventResponseDTO(&event)
		responses = append(responses, *response)
	}

	return c.Status(fiber.StatusOK).JSON(responses)
}

// UpdateEvent godoc
// @Summary Update an event
// @Description Update an existing event in the system
// @Tags Event
// @Accept json
// @Produce json
// @Param id path int true "Event ID"
// @Param event body dto.EventCreateDTO true "Updated event information"
// @Success 200 {object} dto.EventResponseDTO "Successfully updated event"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 404 {object} map[string]interface{} "Event not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /event/{id} [put]
func (h *HttpEventHandler) UpdateEvent(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid event ID",
		})
	}

	// ตรวจสอบว่า event มีอยู่จริงหรือไม่
	existingEvent, err := h.eventUseCase.GetEventByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Event not found",
		})
	}

	// Parse ข้อมูลที่ต้องการอัปเดต
	var input dto.EventCreateDTO
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// อัปเดตข้อมูล event จาก input
	existingEvent.Name = input.Name
	existingEvent.Year = input.Year
	existingEvent.Rationale = input.Rationale
	existingEvent.Objective = input.Objective
	existingEvent.StartDate = input.StartDate
	existingEvent.EndDate = input.EndDate
	existingEvent.Location = input.Location
	existingEvent.Methodology = input.Methodology
	existingEvent.HasBudget = input.HasBudget
	existingEvent.Monitoring = input.Monitoring
	existingEvent.EventTypeStatusID = input.EventTypeStatusID
	existingEvent.EventPlanID = input.EventPlanID
	existingEvent.EventTypeID = input.EventTypeID
	existingEvent.EventStrategyID = input.EventStrategyID
	existingEvent.EventStatusID = input.EventStatusID

	// บันทึกการอัปเดต
	if err := h.eventUseCase.UpdateEvent(existingEvent); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	response := dto.ToEventResponseDTO(existingEvent)
	return c.Status(fiber.StatusOK).JSON(response)
}

// DeleteEvent godoc
// @Summary Delete an event
// @Description Delete an event from the system
// @Tags Event
// @Accept json
// @Produce json
// @Param id path int true "Event ID"
// @Success 200 {object} map[string]interface{} "Successfully deleted event"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 404 {object} map[string]interface{} "Event not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /event/{id} [delete]
func (h *HttpEventHandler) DeleteEvent(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid event ID",
		})
	}

	// ตรวจสอบว่า event มีอยู่จริงหรือไม่
	_, err = h.eventUseCase.GetEventByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Event not found",
		})
	}

	// ลบ event
	if err := h.eventUseCase.DeleteEvent(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Event deleted successfully",
	})
}

// GetCalendarEvents godoc
// @Summary Get all calendar events
// @Description Retrieve all events and return them ordered by start date
// @Tags Calendar
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} dto.CalendarResponseDTO "Successfully retrieved all calendar events"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /events/calendar [get]
func (h *HttpEventHandler) GetCalendarEvents(c *fiber.Ctx) error {
	events, err := h.eventUseCase.GetAllEvents()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].StartDate.Before(events[j].StartDate)
	})

	var calendarResponses []dto.CalendarResponseDTO
	for _, event := range events {
		response := dto.ToCalendarResponseDTO(&event)
		calendarResponses = append(calendarResponses, response)
	}

	// ส่งผลลัพธ์กลับไปยังผู้ใช้
	return c.Status(fiber.StatusOK).JSON(calendarResponses)
}

// UpdateEventStatus godoc
// @Summary Update event status
// @Description Update only the status of an existing event, if status is 4, notify all users
// @Tags Event
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Event ID"
// @Param status body dto.EventStatusUpdateDTO true "New event status"
// @Success 200 {object} dto.EventResponseDTO "Successfully updated event status"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 404 {object} map[string]interface{} "Event not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /event/{id}/status [put]
func (h *HttpEventHandler) UpdateEventStatus(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid event ID",
		})
	}

	// ตรวจสอบว่า event มีอยู่จริงหรือไม่
	existingEvent, err := h.eventUseCase.GetEventByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Event not found",
		})
	}

	var input dto.EventStatusUpdateDTO
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if input.EventStatusID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "event_status_id is required",
		})
	}

	// อัพเดทเฉพาะ status
	existingEvent.EventStatusID = input.EventStatusID

	if err := h.eventUseCase.UpdateEvent(existingEvent); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// ถ้า status = 4 สร้าง notification ให้ทุก user
	if input.EventStatusID == 4 {
		users, err := h.userUseCase.GetAllUsers() // ใช้ UserUseCase แทน EventUseCase
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to get users: " + err.Error(),
			})
		}

		for _, user := range users {
			notification := entities.Notification{
				UserID:  user.UserID,
				EventID: uint(id),
				Active:  true,
			}
			if err := h.notifyUseCase.CreateNotify(&notification); err != nil {
				// Log error but continue with other users
				continue
			}
		}
	}

	response := dto.ToEventResponseDTO(existingEvent)
	return c.Status(fiber.StatusOK).JSON(response)
}
