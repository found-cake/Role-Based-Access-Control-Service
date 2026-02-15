package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"role-based-access-control-service/config"
	"role-based-access-control-service/db"
	"role-based-access-control-service/dto"
	"role-based-access-control-service/handlers"
	appmiddleware "role-based-access-control-service/middleware"
	"role-based-access-control-service/pkg/httpx"
	"role-based-access-control-service/service"
	"role-based-access-control-service/validation"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.Load()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	database, err := db.Connect(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	if err := database.AutoMigrate(); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	defer func() {
		if err := database.Close(); err != nil {
			log.Printf("database close error: %v", err)
		}
	}()

	userRepository := db.NewUserRepository(database.Conn)
	authService := service.NewAuthService(userRepository, cfg.JWTSecret)
	authHandler := handlers.NewAuthHandler(authService)
	authMiddleware := appmiddleware.NewAuthMiddleware(cfg.JWTSecret)

	e := echo.New()
	e.HideBanner = true
	e.Validator = validation.New()
	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.Recover())

	e.POST("/auth/register", authHandler.Register)
	e.POST("/auth/login", authHandler.Login)
	e.GET("/auth/me", authHandler.Me, authMiddleware.RequireAuth())
	e.GET("/health", func(c echo.Context) error {
		return httpx.JSON(c, http.StatusOK, dto.APIResponse{Success: true, Data: map[string]string{"status": "ok"}})
	})

	e.Static("/", "public")

	go func() {
		log.Printf("server is running on port %s", cfg.Port)
		if err := e.Start(":" + cfg.Port); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server failed: %v", err)
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()
	if err := e.Shutdown(shutdownCtx); err != nil {
		log.Printf("graceful shutdown failed: %v", err)
	}
}
