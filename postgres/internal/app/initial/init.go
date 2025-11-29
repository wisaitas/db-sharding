package initial

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/wisaitas/db-sharding/pkg/db/postgres"
	"github.com/wisaitas/db-sharding/postgres/internal/app"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("warning: failed to load .env file: %v", err)
	}

	if err := env.Parse(&app.Config); err != nil {
		log.Fatalf("error parsing config: %v", err)
	}
}

type App struct {
	client   *client
	fiberApp *fiber.App
}

func NewApp() *App {
	client := newClient()

	fiberApp := fiber.New()

	fiberApp.Use(logger.New())

	fiberApp.Get("/hello-world", func(c fiber.Ctx) error {

		return c.JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})

	return &App{
		client:   client,
		fiberApp: fiberApp,
	}
}

func (i *App) Start() error {
	return i.fiberApp.Listen(fmt.Sprintf(":%s", app.Config.Server.Port))
}

func (i *App) Stop() {
	if err := i.fiberApp.Shutdown(); err != nil {
		log.Fatalf("error stopping server: %v", err)
	}

	if err := postgres.Close(i.client.postgresShard1); err != nil {
		log.Fatal(err)
	}

	if err := postgres.Close(i.client.postgresShard2); err != nil {
		log.Fatal(err)
	}

	log.Println("server stopped")
}
