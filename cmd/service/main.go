package main

import "fmt"

// App contains all the system depedencies like database connection, etc.
type App struct {
}

func (app *App) Run() error {
	fmt.Println("Setting up out APP")
	return nil
}

func main() {
	fmt.Print("hello")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up out REST API")
		fmt.Println(err)
	}
}
