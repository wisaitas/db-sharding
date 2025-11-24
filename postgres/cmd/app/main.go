package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/wisaitas/db-sharding/postgres/internal/app/initial"
)

func main() {
	app := initial.NewApp()

	go func() {
		if err := app.Start(); err != nil {
			log.Fatalf("error starting server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	app.Stop()
}
