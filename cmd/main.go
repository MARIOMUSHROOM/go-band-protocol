package main

import (
	httpHandler "band_protocol_go/internal/adaptors/http"
	"band_protocol_go/internal/application"
	"band_protocol_go/pkg/client"
	"band_protocol_go/pkg/config"
	"band_protocol_go/pkg/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	// Setup Fiber app
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		// log.Info(fmt.Sprintf("Request Body: %s", string(c.Body())))
		return c.Next()
	})
	// Initialize client
	cfg := config.LoadConfig()
	apiClient := client.NewClient(cfg)

	// Initialize services
	actionService := application.NewActionService()
	clientService := service.NewAPIService(apiClient)

	// Initialize HTTP handlers
	actionHandler := httpHandler.NewActionHandler(actionService, clientService)

	publicApi := app.Group("/public")
	publicApi.Get("/boss-baby/:value", actionHandler.GetBossBaby)
	publicApi.Post("/superman-chicken", actionHandler.SupermanChicken)
	publicApi.Post("/transaction", actionHandler.Transaction)

	// Start Fiber server
	err := app.Listen(":8080")
	if err != nil {
		log.Fatal("Fiber server start")
	}
}
