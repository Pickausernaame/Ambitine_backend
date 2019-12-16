package main

import (
	"fmt"

	"github.com/Pickausernaame/Ambitine_backend/server"
)

func main() {
	app := server.New()

	if app == nil {
		fmt.Println("Unable to create server")
	}

	err := app.InitDB("postgres://ambitine:1488@localhost:5432")

	if err != nil {
		fmt.Println("Unable to init database:", err)
	}

	app.Run(":9090")
}
