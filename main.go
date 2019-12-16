package main

import (
	"fmt"
	"os"

	"github.com/Pickausernaame/Ambitine_backend/server"
)

func main() {
	// Создание приложения
	app := server.New()
	if app == nil {
		fmt.Println("Unable to create app")
		os.Exit(1)
	}
	app.Run(":9090")
}
