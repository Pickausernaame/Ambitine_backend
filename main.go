package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Pickausernaame/Ambitine_backend/server"
)

func main() {
	// todo написать ручки для обещаний
	// todo связать ручки обещаний с бд
	// todo рефакторинг кода

	// Пауза для полноценного запуска контейнера бд
	time.Sleep(5 * time.Second)
	// Создание приложения
	app := server.New()
	if app == nil {
		fmt.Println("Unable to create app")
		os.Exit(1)
	}
	app.Run(":9090")
}
