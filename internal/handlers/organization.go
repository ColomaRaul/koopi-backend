package handlers

import (
	"net/http"

	"koopi-backend/internal/database"
	"koopi-backend/pkg/models"

	"github.com/gin-gonic/gin"
)

type OrganizationHandler struct {
	db *database.DB
}

func NewOrganizationHandler(db *database.DB) *OrganizationHandler {
	return &OrganizationHandler{db: db}
}

// GetOrganizations obtiene la lista de organizaciones
func (h *OrganizationHandler) GetOrganizations(c *gin.Context) {
	var organizations []models.Organization
	
	if err := h.db.Preload("Owner").Find(&organizations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Error al obtener organizaciones",
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    organizations,
		"message": "Organizaciones obtenidas exitosamente",
	})
}

// GetOrganization obtiene una organización específica
func (h *OrganizationHandler) GetOrganization(c *gin.Context) {
	id := c.Param("id")
	
	var organization models.Organization
	if err := h.db.Preload("Owner").First(&organization, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Organización no encontrada",
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    organization,
		"message": "Organización obtenida exitosamente",
	})
}

// CreateOrganization crea una nueva organización
func (h *OrganizationHandler) CreateOrganization(c *gin.Context) {
	var input struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
		Website     string `json:"website"`
		Phone       string `json:"phone"`
		Email       string `json:"email"`
	}
	
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Datos inválidos: " + err.Error(),
		})
		return
	}
	
	userID, _ := c.Get("user_id")
	
	organization := models.Organization{
		Name:        input.Name,
		Description: input.Description,
		OwnerID:     userID.(uint),
		Website:     input.Website,
		Phone:       input.Phone,
		Email:       input.Email,
		IsVerified:  false,
	}
	
	if err := h.db.Create(&organization).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Error al crear la organización",
		})
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    organization,
		"message": "Organización creada exitosamente",
	})
}

// UpdateOrganization actualiza una organización
func (h *OrganizationHandler) UpdateOrganization(c *gin.Context) {
	id := c.Param("id")
	
	var organization models.Organization
	if err := h.db.First(&organization, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Organización no encontrada",
		})
		return
	}
	
	// Verificar que el usuario sea el propietario
	userID, _ := c.Get("user_id")
	if organization.OwnerID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"error":   "No tienes permisos para editar esta organización",
		})
		return
	}
	
	var input struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Website     string `json:"website"`
		Phone       string `json:"phone"`
		Email       string `json:"email"`
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
		organization.Name = input.Name
	}
	if input.Description != "" {
		organization.Description = input.Description
	}
	if input.Website != "" {
		organization.Website = input.Website
	}
	if input.Phone != "" {
		organization.Phone = input.Phone
	}
	if input.Email != "" {
		organization.Email = input.Email
	}
	
	if err := h.db.Save(&organization).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Error al actualizar la organización",
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    organization,
		"message": "Organización actualizada exitosamente",
	})
}

// DeleteOrganization elimina una organización
func (h *OrganizationHandler) DeleteOrganization(c *gin.Context) {
	id := c.Param("id")
	
	var organization models.Organization
	if err := h.db.First(&organization, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Organización no encontrada",
		})
		return
	}
	
	// Verificar que el usuario sea el propietario
	userID, _ := c.Get("user_id")
	if organization.OwnerID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"error":   "No tienes permisos para eliminar esta organización",
		})
		return
	}
	
	if err := h.db.Delete(&organization).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Error al eliminar la organización",
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Organización eliminada exitosamente",
	})
}

// GetOrganizationEvents obtiene los eventos de una organización
func (h *OrganizationHandler) GetOrganizationEvents(c *gin.Context) {
	id := c.Param("id")
	
	var events []models.Event
	if err := h.db.Preload("Category").Preload("Location").
		Where("organizer_id = ?", id).Find(&events).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Error al obtener eventos",
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    events,
		"message": "Eventos de la organización obtenidos exitosamente",
	})
} 