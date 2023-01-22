package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/TonyPath/go-rest-api/internal/transport/http"
)

// App contains all the system depedencies like database connection, etc.
type App struct {
}

func (app *App) Run() error {
	fmt.Println("Setting up out APP")

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to setup server")
		return err
	}

	return nil
}

func main() {
	fmt.Println("Go REST API")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up out REST API")
		fmt.Println(err)
	}
}
