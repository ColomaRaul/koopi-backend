package handlers

import (
	"net/http"

	"koopi-backend/internal/database"
	"koopi-backend/pkg/models"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	db *database.DB
}

func NewCategoryHandler(db *database.DB) *CategoryHandler {
	return &CategoryHandler{db: db}
}

// GetCategories obtiene la lista de categorías
func (h *CategoryHandler) GetCategories(c *gin.Context) {
	var categories []models.EventCategory
	
	if err := h.db.Where("is_active = ?", true).Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Error al obtener categorías",
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    categories,
		"message": "Categorías obtenidas exitosamente",
	})
}

// CreateCategory crea una nueva categoría (solo admin)
func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var input struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
		Color       string `json:"color"`
	}
	
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Datos inválidos: " + err.Error(),
		})
		return
	}
	
	category := models.EventCategory{
		Name:        input.Name,
		Description: input.Description,
		Icon:        input.Icon,
		Color:       input.Color,
		IsActive:    true,
	}
	
	if err := h.db.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Error al crear la categoría",
		})
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    category,
		"message": "Categoría creada exitosamente",
	})
}

// UpdateCategory actualiza una categoría (solo admin)
func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	
	var category models.EventCategory
	if err := h.db.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Categoría no encontrada",
		})
		return
	}
	
	var input struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
		Color       string `json:"color"`
		IsActive    *bool  `json:"is_active"`
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
		category.Name = input.Name
	}
	if input.Description != "" {
		category.Description = input.Description
	}
	if input.Icon != "" {
		category.Icon = input.Icon
	}
	if input.Color != "" {
		category.Color = input.Color
	}
	if input.IsActive != nil {
		category.IsActive = *input.IsActive
	}
	
	if err := h.db.Save(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Error al actualizar la categoría",
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    category,
		"message": "Categoría actualizada exitosamente",
	})
}

// DeleteCategory elimina una categoría (solo admin)
func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	
	var category models.EventCategory
	if err := h.db.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Categoría no encontrada",
		})
		return
	}
	
	if err := h.db.Delete(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Error al eliminar la categoría",
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Categoría eliminada exitosamente",
	})
} 