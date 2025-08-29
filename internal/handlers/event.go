package handlers

import (
	"net/http"
	"time"

	"koopi-backend/internal/database"
	"koopi-backend/pkg/models"

	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	db *database.DB
}

func NewEventHandler(db *database.DB) *EventHandler {
	return &EventHandler{db: db}
}

// GetEvents obtiene la lista de eventos
func (h *EventHandler) GetEvents(c *gin.Context) {
	var events []models.Event
	
	query := h.db.Preload("Category").Preload("Organizer").Preload("Location")
	
	// Aplicar filtros
	if categoryID := c.Query("category_id"); categoryID != "" {
		query = query.Where("category_id = ?", categoryID)
	}
	
	if city := c.Query("city"); city != "" {
		query = query.Joins("JOIN locations ON events.location_id = locations.id").
			Where("locations.city ILIKE ?", "%"+city+"%")
	}
	
	// Solo eventos activos
	query = query.Where("is_active = ?", true)
	
	if err := query.Find(&events).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Error al obtener eventos",
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    events,
		"message": "Eventos obtenidos exitosamente",
	})
}

// GetEvent obtiene un evento específico
func (h *EventHandler) GetEvent(c *gin.Context) {
	id := c.Param("id")
	
	var event models.Event
	if err := h.db.Preload("Category").Preload("Organizer").Preload("Location").
		First(&event, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Evento no encontrado",
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    event,
		"message": "Evento obtenido exitosamente",
	})
}

// CreateEvent crea un nuevo evento
func (h *EventHandler) CreateEvent(c *gin.Context) {
	var input struct {
		Title       string  `json:"title" binding:"required"`
		Description string  `json:"description"`
		CategoryID  uint    `json:"category_id" binding:"required"`
		LocationID  uint    `json:"location_id" binding:"required"`
		StartDate   string  `json:"start_date" binding:"required"`
		EndDate     string  `json:"end_date" binding:"required"`
		Price       float64 `json:"price"`
		Capacity    int     `json:"capacity"`
		ImageURL    string  `json:"image_url"`
		Website     string  `json:"website"`
	}
	
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Datos inválidos: " + err.Error(),
		})
		return
	}
	
	// Obtener el organizador del usuario autenticado
	userID, _ := c.Get("user_id")
	
	// Buscar la organización del usuario
	var organization models.Organization
	if err := h.db.Where("owner_id = ?", userID).First(&organization).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Debes tener una organización para crear eventos",
		})
		return
	}
	
	// Parsear fechas (implementación básica)
	startDate := time.Now() // Por ahora usamos tiempo actual
	endDate := time.Now().Add(2 * time.Hour) // Por ahora 2 horas después
	
	event := models.Event{
		Title:       input.Title,
		Description: input.Description,
		CategoryID:  input.CategoryID,
		OrganizerID: organization.ID,
		LocationID:  input.LocationID,
		StartDate:   startDate,
		EndDate:     endDate,
		Price:       input.Price,
		Capacity:    input.Capacity,
		ImageURL:    input.ImageURL,
		Website:     input.Website,
		IsActive:    true,
	}
	
	if err := h.db.Create(&event).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Error al crear el evento",
		})
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    event,
		"message": "Evento creado exitosamente",
	})
}

// UpdateEvent actualiza un evento
func (h *EventHandler) UpdateEvent(c *gin.Context) {
	id := c.Param("id")
	
	var event models.Event
	if err := h.db.First(&event, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Evento no encontrado",
		})
		return
	}
	
	// Verificar que el usuario sea el organizador
	userID, _ := c.Get("user_id")
	if event.OrganizerID != uint(userID.(uint)) {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"error":   "No tienes permisos para editar este evento",
		})
		return
	}
	
	var input struct {
		Title       string  `json:"title"`
		Description string  `json:"description"`
		CategoryID  uint    `json:"category_id"`
		LocationID  uint    `json:"location_id"`
		StartDate   string  `json:"start_date"`
		EndDate     string  `json:"end_date"`
		Price       float64 `json:"price"`
		Capacity    int     `json:"capacity"`
		ImageURL    string  `json:"image_url"`
		Website     string  `json:"website"`
		IsActive    *bool   `json:"is_active"`
	}
	
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Datos inválidos: " + err.Error(),
		})
		return
	}
	
	// Actualizar campos
	if input.Title != "" {
		event.Title = input.Title
	}
	if input.Description != "" {
		event.Description = input.Description
	}
	if input.CategoryID != 0 {
		event.CategoryID = input.CategoryID
	}
	if input.LocationID != 0 {
		event.LocationID = input.LocationID
	}
	if input.StartDate != "" {
		// Por ahora no actualizamos las fechas
		// event.StartDate = startDate
	}
	if input.EndDate != "" {
		// Por ahora no actualizamos las fechas
		// event.EndDate = endDate
	}
	if input.Price >= 0 {
		event.Price = input.Price
	}
	if input.Capacity > 0 {
		event.Capacity = input.Capacity
	}
	if input.ImageURL != "" {
		event.ImageURL = input.ImageURL
	}
	if input.Website != "" {
		event.Website = input.Website
	}
	if input.IsActive != nil {
		event.IsActive = *input.IsActive
	}
	
	if err := h.db.Save(&event).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Error al actualizar el evento",
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    event,
		"message": "Evento actualizado exitosamente",
	})
}

// DeleteEvent elimina un evento
func (h *EventHandler) DeleteEvent(c *gin.Context) {
	id := c.Param("id")
	
	var event models.Event
	if err := h.db.First(&event, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Evento no encontrado",
		})
		return
	}
	
	// Verificar que el usuario sea el organizador
	userID, _ := c.Get("user_id")
	if event.OrganizerID != uint(userID.(uint)) {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"error":   "No tienes permisos para eliminar este evento",
		})
		return
	}
	
	if err := h.db.Delete(&event).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Error al eliminar el evento",
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Evento eliminado exitosamente",
	})
}

// GetNearbyEvents obtiene eventos cercanos
func (h *EventHandler) GetNearbyEvents(c *gin.Context) {
	// Por ahora, simplemente retornamos todos los eventos
	// En el futuro, implementaremos la lógica de geolocalización
	h.GetEvents(c)
}

// SearchEvents busca eventos
func (h *EventHandler) SearchEvents(c *gin.Context) {
	// Por ahora, simplemente retornamos todos los eventos
	// En el futuro, implementaremos la lógica de búsqueda
	h.GetEvents(c)
}

 