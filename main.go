package main

import (
	"context"
	"fmt"

	"github.com/deboraamartinez/orders-api/application"
)

func main() {
	app := application.New(application.LoadConfig())
	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}
