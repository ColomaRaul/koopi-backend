package handlers

import (
	"net/http"

	"koopi-backend/internal/database"
	"koopi-backend/pkg/models"

	"github.com/gin-gonic/gin"
)

type LocationHandler struct {
	db *database.DB
}

func NewLocationHandler(db *database.DB) *LocationHandler {
	return &LocationHandler{db: db}
}

// GetLocations obtiene la lista de ubicaciones
func (h *LocationHandler) GetLocations(c *gin.Context) {
	var locations []models.Location

	if err := h.db.Find(&locations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Error al obtener ubicaciones",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    locations,
		"message": "Ubicaciones obtenidas exitosamente",
	})
}

// GetLocation obtiene una ubicación específica
func (h *LocationHandler) GetLocation(c *gin.Context) {
	id := c.Param("id")

	var location models.Location
	if err := h.db.First(&location, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Ubicación no encontrada",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    location,
		"message": "Ubicación obtenida exitosamente",
	})
}

// CreateLocation crea una nueva ubicación
func (h *LocationHandler) CreateLocation(c *gin.Context) {
	var input struct {
		Name       string  `json:"name" binding:"required"`
		Address    string  `json:"address" binding:"required"`
		City       string  `json:"city" binding:"required"`
		State      string  `json:"state"`
		Country    string  `json:"country"`
		PostalCode string  `json:"postal_code"`
		Latitude   float64 `json:"latitude" binding:"required"`
		Longitude  float64 `json:"longitude" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Datos inválidos: " + err.Error(),
		})
		return
	}

	location := models.Location{
		Name:       input.Name,
		Address:    input.Address,
		City:       input.City,
		State:      input.State,
		Country:    input.Country,
		PostalCode: input.PostalCode,
		Latitude:   input.Latitude,
		Longitude:  input.Longitude,
	}

	if err := h.db.Create(&location).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Error al crear la ubicación",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    location,
		"message": "Ubicación creada exitosamente",
	})
}

// UpdateLocation actualiza una ubicación
func (h *LocationHandler) UpdateLocation(c *gin.Context) {
	id := c.Param("id")

	var location models.Location
	if err := h.db.First(&location, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Ubicación no encontrada",
		})
		return
	}

	var input struct {
		Name       string  `json:"name"`
		Address    string  `json:"address"`
		City       string  `json:"city"`
		State      string  `json:"state"`
		Country    string  `json:"country"`
		PostalCode string  `json:"postal_code"`
		Latitude   float64 `json:"latitude"`
		Longitude  float64 `json:"longitude"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Datos inválidos: " + err.Error(),
		})
		return
	}

	// Actualizar campos
	if input.Name != "" {
		location.Name = input.Name
	}
	if input.Address != "" {
		location.Address = input.Address
	}
	if input.City != "" {
		location.City = input.City
	}
	if input.State != "" {
		location.State = input.State
	}
	if input.Country != "" {
		location.Country = input.Country
	}
	if input.PostalCode != "" {
		location.PostalCode = input.PostalCode
	}
	if input.Latitude != 0 {
		location.Latitude = input.Latitude
	}
	if input.Longitude != 0 {
		location.Longitude = input.Longitude
	}

	if err := h.db.Save(&location).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Error al actualizar la ubicación",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    location,
		"message": "Ubicación actualizada exitosamente",
	})
}

// DeleteLocation elimina una ubicación
func (h *LocationHandler) DeleteLocation(c *gin.Context) {
	id := c.Param("id")

	var location models.Location
	if err := h.db.First(&location, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Ubicación no encontrada",
		})
		return
	}

	if err := h.db.Delete(&location).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Error al eliminar la ubicación",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Ubicación eliminada exitosamente",
	})
}

// SearchLocations busca ubicaciones
func (h *LocationHandler) SearchLocations(c *gin.Context) {
	query := c.Query("q")

		var locations []models.Location
	dbQuery := h.db.GetDB()
	
	if query != "" {
		dbQuery = dbQuery.Where("name ILIKE ? OR address ILIKE ? OR city ILIKE ?", 
			"%"+query+"%", "%"+query+"%", "%"+query+"%")
	}
	
	if err := dbQuery.Find(&locations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Error al buscar ubicaciones",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    locations,
		"message": "Búsqueda de ubicaciones completada",
	})
}
