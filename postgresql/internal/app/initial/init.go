package initial

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
	"github.com/wisaitas/db-sharding/postgresql/internal/app"
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
	fiberApp *fiber.App
}

func NewApp() *App {
	fiberApp := fiber.New()

	// newMiddleware(fiberApp)

	return &App{
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

	log.Println("server stopped")
}
