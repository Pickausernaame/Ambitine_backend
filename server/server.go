package server

import (
	"fmt"

	db "github.com/Pickausernaame/Ambitine_backend/server/db"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
)

type App struct {
	Router *gin.Engine
	DB     *db.DBHandler
}

func (instance *App) initializeRoutes() {
	api := instance.Router.Group("/api")
	{
		api.GET("/hello", HelloFunc)
		api.POST("/signin", instance.SignInHand)
		api.POST("/signup", instance.SignUpHand)
	}
}

func New() *App {
	a := App{Router: gin.New()}

	a.initializeRoutes()

	a.Router.Use(gin.Recovery())
	a.Router.Use(gin.Logger())

	return &a
}

func (instance *App) InitDB(dbCfg string) (err error) {
	var (
		pgxCfg pgx.ConnConfig
	)

	pgxCfg, err = pgx.ParseURI(dbCfg)

	if err != nil {
		fmt.Println("Unable to parse database config:", err)
		return
	}

	conn, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig:     pgxCfg,
		MaxConnections: 128,
	})

	if err != nil {
		fmt.Println("Unbable to build new database connection pool:", err)
		return
	}

	instance.DB = &db.DBHandler{Connection: conn}

	err = instance.DB.ResetDB()

	if err != nil {
		fmt.Println("Unable to create database tables:", err)
	}

	return
}

func (instance *App) Run(port string) {
	err := instance.Router.Run(port)
	fmt.Print(err)
}
