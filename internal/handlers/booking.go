package handlers

import (
	"net/http"

	"koopi-backend/internal/database"
	"koopi-backend/pkg/models"

	"github.com/gin-gonic/gin"
)

type BookingHandler struct {
	db *database.DB
}

func NewBookingHandler(db *database.DB) *BookingHandler {
	return &BookingHandler{db: db}
}

// GetUserBookings obtiene las reservas del usuario autenticado
func (h *BookingHandler) GetUserBookings(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var bookings []models.Booking
	if err := h.db.Preload("Event").Preload("Event.Category").
		Where("user_id = ?", userID).Find(&bookings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Error al obtener reservas",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    bookings,
		"message": "Reservas obtenidas exitosamente",
	})
}

// CreateBooking crea una nueva reserva
func (h *BookingHandler) CreateBooking(c *gin.Context) {
	var input struct {
		EventID    uint    `json:"event_id" binding:"required"`
		Quantity   int     `json:"quantity"`
		TotalPrice float64 `json:"total_price"`
		Notes      string  `json:"notes"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Datos inválidos: " + err.Error(),
		})
		return
	}

	userID, _ := c.Get("user_id")

	// Verificar que el evento existe y está activo
	var event models.Event
	if err := h.db.First(&event, input.EventID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Evento no encontrado",
		})
		return
	}

	if !event.IsActive {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "El evento no está activo",
		})
		return
	}

	// Verificar que no haya una reserva duplicada
	var existingBooking models.Booking
	if err := h.db.Where("event_id = ? AND user_id = ?", input.EventID, userID).
		First(&existingBooking).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Ya tienes una reserva para este evento",
		})
		return
	}

	booking := models.Booking{
		EventID:    input.EventID,
		UserID:     userID.(uint),
		Quantity:   input.Quantity,
		TotalPrice: input.TotalPrice,
		Notes:      input.Notes,
		Status:     "pending",
	}

	if err := h.db.Create(&booking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Error al crear la reserva",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    booking,
		"message": "Reserva creada exitosamente",
	})
}

// GetBooking obtiene una reserva específica
func (h *BookingHandler) GetBooking(c *gin.Context) {
	id := c.Param("id")
	userID, _ := c.Get("user_id")

	var booking models.Booking
	if err := h.db.Preload("Event").Preload("Event.Category").
		Where("id = ? AND user_id = ?", id, userID).First(&booking).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Reserva no encontrada",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    booking,
		"message": "Reserva obtenida exitosamente",
	})
}

// CancelBooking cancela una reserva
func (h *BookingHandler) CancelBooking(c *gin.Context) {
	id := c.Param("id")
	userID, _ := c.Get("user_id")

	var booking models.Booking
	if err := h.db.Where("id = ? AND user_id = ?", id, userID).First(&booking).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Reserva no encontrada",
		})
		return
	}

	if booking.Status == "cancelled" {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "La reserva ya está cancelada",
		})
		return
	}

	booking.Status = "cancelled"
	if err := h.db.Save(&booking).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Error al cancelar la reserva",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    booking,
		"message": "Reserva cancelada exitosamente",
	})
}

// GetEventBookings obtiene las reservas de un evento (solo organizador)
func (h *BookingHandler) GetEventBookings(c *gin.Context) {
	eventID := c.Param("id")
	userID, _ := c.Get("user_id")

	// Verificar que el usuario sea el organizador del evento
	var event models.Event
	if err := h.db.Where("id = ?", eventID).First(&event).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Evento no encontrado",
		})
		return
	}

	// Buscar la organización del usuario
	var organization models.Organization
	if err := h.db.Where("owner_id = ?", userID).First(&organization).Error; err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"error":   "No tienes permisos para ver las reservas de este evento",
		})
		return
	}

	if event.OrganizerID != organization.ID {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"error":   "No tienes permisos para ver las reservas de este evento",
		})
		return
	}

	var bookings []models.Booking
	if err := h.db.Preload("User").Where("event_id = ?", eventID).Find(&bookings).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Error al obtener reservas",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    bookings,
		"message": "Reservas del evento obtenidas exitosamente",
	})
}
