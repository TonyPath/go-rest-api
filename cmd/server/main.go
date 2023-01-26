package main

import (
	"fmt"

	"github.com/TonyPath/go-rest-api/internal/comment"
	"github.com/TonyPath/go-rest-api/internal/db"
	transportHTTP "github.com/TonyPath/go-rest-api/internal/transport/http"
)

func main() {
	fmt.Println("Go REST API")

	if err := run(); err != nil {
		fmt.Println("Error starting up out REST API")
		fmt.Println(err)
	}
}

func run() error {
	fmt.Println("starting up our application")

	db, err := db.NewDatabase()
	if err != nil {
		return err
	}

	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}

	cmtService := comment.NewService(db)

	httpHandler := transportHTTP.NewHandler(cmtService)
	if err := httpHandler.Serve(); err != nil {
		return nil
	}

	return nil
}
