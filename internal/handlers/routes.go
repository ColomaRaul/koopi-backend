package handlers

import (
	"koopi-backend/internal/database"

	"github.com/gin-gonic/gin"
)

// SetupPublicRoutes configura las rutas públicas
func SetupPublicRoutes(router *gin.RouterGroup, db *database.DB) {
	// Auth routes
	authHandler := NewAuthHandler(db)
	router.POST("/auth/register", authHandler.Register)
	router.POST("/auth/login", authHandler.Login)

	// Event routes (públicas)
	eventHandler := NewEventHandler(db)
	router.GET("/events", eventHandler.GetEvents)
	router.GET("/events/:id", eventHandler.GetEvent)
	router.GET("/events/nearby", eventHandler.GetNearbyEvents)
	router.GET("/events/search", eventHandler.SearchEvents)

	// Category routes
	categoryHandler := NewCategoryHandler(db)
	router.GET("/categories", categoryHandler.GetCategories)

	// Location routes
	locationHandler := NewLocationHandler(db)
	router.GET("/locations", locationHandler.GetLocations)
	router.GET("/locations/:id", locationHandler.GetLocation)
	router.GET("/locations/search", locationHandler.SearchLocations)

	// Organization routes (públicas)
	orgHandler := NewOrganizationHandler(db)
	router.GET("/organizations", orgHandler.GetOrganizations)
	router.GET("/organizations/:id", orgHandler.GetOrganization)
	router.GET("/organizations/:id/events", orgHandler.GetOrganizationEvents)
}

// SetupProtectedRoutes configura las rutas protegidas
func SetupProtectedRoutes(router *gin.RouterGroup, db *database.DB) {
	// Auth routes
	authHandler := NewAuthHandler(db)
	router.GET("/auth/profile", authHandler.GetProfile)
	router.PUT("/auth/profile", authHandler.UpdateProfile)

	// Event routes (protegidas)
	eventHandler := NewEventHandler(db)
	router.POST("/events", eventHandler.CreateEvent)
	router.PUT("/events/:id", eventHandler.UpdateEvent)
	router.DELETE("/events/:id", eventHandler.DeleteEvent)

	// Booking routes
	bookingHandler := NewBookingHandler(db)
	router.GET("/bookings", bookingHandler.GetUserBookings)
	router.POST("/bookings", bookingHandler.CreateBooking)
	router.GET("/bookings/:id", bookingHandler.GetBooking)
	router.PUT("/bookings/:id/cancel", bookingHandler.CancelBooking)
	router.GET("/events/:id/bookings", bookingHandler.GetEventBookings)

	// Organization routes (protegidas)
	orgHandler := NewOrganizationHandler(db)
	router.POST("/organizations", orgHandler.CreateOrganization)
	router.PUT("/organizations/:id", orgHandler.UpdateOrganization)
	router.DELETE("/organizations/:id", orgHandler.DeleteOrganization)

	// Location routes (protegidas)
	locationHandler := NewLocationHandler(db)
	router.POST("/locations", locationHandler.CreateLocation)
	router.PUT("/locations/:id", locationHandler.UpdateLocation)
	router.DELETE("/locations/:id", locationHandler.DeleteLocation)

	// Category routes (protegidas - solo admin)
	categoryHandler := NewCategoryHandler(db)
	router.POST("/categories", categoryHandler.CreateCategory)
	router.PUT("/categories/:id", categoryHandler.UpdateCategory)
	router.DELETE("/categories/:id", categoryHandler.DeleteCategory)
} 