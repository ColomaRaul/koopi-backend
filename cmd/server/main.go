package main

import (
	"log"
	"os"

	"koopi-backend/internal/config"
	"koopi-backend/internal/database"
	"koopi-backend/internal/handlers"
	"koopi-backend/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "koopi-backend/docs"
)

// @title           Koopi Events API
// @version         1.0
// @description     API para gestión de eventos
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Configurar la aplicación
	cfg := config.Load()

	// Inicializar base de datos
	db, err := database.Init(cfg.Database)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Configurar Gin
	if cfg.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Crear router
	router := gin.Default()

	// Middleware global
	router.Use(middleware.CORS())
	router.Use(middleware.Logger())

	// Configurar rutas
	setupRoutes(router, db)

	// Iniciar servidor
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func setupRoutes(router *gin.Engine, db *database.DB) {
	// API v1
	v1 := router.Group("/api/v1")

	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Rutas públicas
	handlers.SetupPublicRoutes(v1, db)

	// Rutas protegidas
	protected := v1.Group("/")
	protected.Use(middleware.Auth())
	handlers.SetupProtectedRoutes(protected, db)

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"message": "Koopi Events API is running",
		})
	})
} 