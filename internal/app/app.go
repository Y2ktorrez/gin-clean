package app

import (
	"fmt"
	"gin-clean/config"
	v1 "gin-clean/internal/controller/http/v1"
	"gin-clean/internal/entity"
	"gin-clean/internal/usecase/repositories"
	"gin-clean/internal/usecase/services"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	router *gin.Engine
	db     *gorm.DB
}

func New(config *config.Config) (*App, error) {
	db, err := setupDatabase(config)
	if err != nil {
		return nil, fmt.Errorf("failed to setup database: %w", err)
	}

	app := &App{
		router: gin.Default(),
		db:     db,
	}

	app.setupRoutes()
	return app, nil
}

func (a *App) Run(addr string) error {
	return a.router.Run(addr)
}

func setupDatabase(config *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.GetDBURL()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto-migrate the database
	if err := db.AutoMigrate(&entity.User{}); err != nil {
		return nil, err
	}

	return db, nil
}

func (a *App) setupRoutes() {
	// Initialize repositories
	userRepo := repositories.NewUserRepository(a.db)

	// Initialize services
	userService := services.NewUserService(userRepo)

	// Setup routes
	v1.SetupRoutes(a.router, userService)
}
