package main

import (
	"fmt"
	"github.com/orzzet/ropero-solidario-api/internal/database"
	"github.com/orzzet/ropero-solidario-api/internal/models"
	transportHTTP "github.com/orzzet/ropero-solidario-api/internal/transport/http"
	"net/http"
)

// App - the struct which contains things like pointers
// to database connections
type App struct {
}

// Run - sets up our application
func (app *App) Run() error {
	fmt.Println("Setting Up Our App")

	var err error
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}
	err = database.MigrateDB(db)
	if err != nil {
		return err
	}
	modelsService := models.NewService(db)

	handler := transportHTTP.NewHandler(modelsService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8850", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}
	fmt.Println("Running server on port 8850")
	return nil
}

func main() {
	fmt.Println("Go REST API Course")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our REST API")
		fmt.Println(err)
	}
}
