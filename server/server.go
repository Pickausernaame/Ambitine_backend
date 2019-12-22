package server

import (
	"errors"
	"fmt"
	"github.com/Pickausernaame/Ambitine_backend/server/kanzler"
	"log"
	"os"
	"strconv"

	"github.com/Pickausernaame/Ambitine_backend/server/middleware"

	db "github.com/Pickausernaame/Ambitine_backend/server/db"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
)

const (
	LOGIN   = "P_LOGIN"
	PASS    = "P_PASSWORD"
	DB_NAME = "P_DB"
	IP      = "P_HOST"
	PORT    = "P_PORT"
)

type App struct {
	Router *gin.Engine
	DB     *db.DBHandler
	WM     *kanzler.WalletManager
}

func (instance *App) initializeRoutes() {

	instance.Router.Use(gin.Recovery())
	instance.Router.Use(gin.Logger())

	api := instance.Router.Group("/api")
	{
		// api.GET("/hello", instance.HelloFunc)

		api.GET("/get_export_promises", middleware.AuthMiddleware(instance.GetAuthorPromises))
		api.GET("/get_import_promises", middleware.AuthMiddleware(instance.GetReceiverPromises))
		api.GET("/user_info", middleware.AuthMiddleware(instance.UserInfo))

		api.POST("/set_new_promis", instance.CreateNewPromise)

		api.GET("/users_autocomplete", middleware.AuthMiddleware(instance.GetAllUsers))

		api.POST("/signin", instance.SignInHand)
		api.POST("/signup", instance.SignUpHand)

		api.POST("/logout", middleware.AuthMiddleware(instance.Logout))
		api.POST("/new_promise", middleware.AuthMiddleware(instance.CreateNewPromise))

	}
}

func New() *App {
	a := App{Router: gin.New()}

	a.initializeRoutes()

	a.Router.Use(gin.Recovery())
	a.Router.Use(gin.Logger())
	err := a.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	a.WM, err = kanzler.New()
	if err != nil {
		log.Fatal(err)
	}

	return &a
}

func dbParamsGetter() (settings *pgx.ConnConfig, err error) {

	settings = &pgx.ConnConfig{}
	exist := false

	settings.User, exist = os.LookupEnv(LOGIN)
	if !exist {
		err = errors.New("Login is not exist in env var: " + LOGIN)
		return nil, err
	}

	settings.Password, exist = os.LookupEnv(PASS)
	if !exist {
		err = errors.New("Login is not exist in env var: " + PASS)
		return nil, err
	}

	settings.Database, exist = os.LookupEnv(DB_NAME)
	if !exist {
		err = errors.New("Login is not exist in env var: " + DB_NAME)
		return nil, err
	}

	settings.Host, exist = os.LookupEnv(IP)
	if !exist {
		err = errors.New("Login is not exist in env var: " + IP)
		return nil, err
	}
	port := ""
	port, exist = os.LookupEnv(PORT)
	if !exist {
		err = errors.New("Login is not exist in env var: " + PORT)
		return nil, err
	}

	intPort, _ := strconv.Atoi(port)
	settings.Port = uint16(intPort)
	return settings, nil
}

func (instance *App) InitDB() (err error) {

	pgxCfg, err := dbParamsGetter()
	if err != nil {
		fmt.Println(err)
		return err
	}

	conn, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig:     *pgxCfg,
		MaxConnections: 128,
	})

	if err != nil {
		fmt.Println("Unbable to build new database connection pool:", err)
		return err
	}

	instance.DB = &db.DBHandler{Connection: conn}

	fmt.Println("Before reset")

	// ##################################################
	{
		err = instance.DB.ResetDB()
		if err != nil {
			fmt.Println("Unable to create database tables:", err)
			return err
		}

		fmt.Println("Before mock")
		mocker := db.Mocker{DB: instance.DB, WM: instance.WM}
		mocker.Mock()
	}
	// ##################################################
	fmt.Println("After reset")
	return nil
}

func (instance *App) Run(port string) {
	err := instance.Router.Run(port)
	fmt.Print(err)
}
