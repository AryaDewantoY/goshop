package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aryadewantoy/goshop/app/databases/seeder"
	"github.com/aryadewantoy/goshop/app/models"
	"github.com/gorilla/mux"
	"github.com/urfave/cli"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	DB        *gorm.DB
	Router    *mux.Router
	AppConfig *AppConfig
}

type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
	AppURL  string
}


type DBConfig struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	DBDriver   string
}

func (routes *Server) Initialize(appConfig AppConfig, dbConfig DBConfig) {
	fmt.Println("Selamat Datang di " + appConfig.AppName)


	routes.initializeDB(dbConfig)
	// routes.initializeAppConfig(appConfig)
	routes.initializeRoutes()
}

func (routes *Server) Run(addr string) {
	fmt.Printf("Listen port %s", addr)
	log.Fatal(http.ListenAndServe(addr, routes.Router))
}

func (routes *Server) initializeDB(dbConfig DBConfig) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		dbConfig.DBHost, dbConfig.DBUser, dbConfig.DBPassword, dbConfig.DBName, dbConfig.DBPort)
	routes.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed on connect to Database!")
	}
}

func (routes *Server) dbMigrate() {
	for _, model := range models.RegisterModels() {
		err := routes.DB.Debug().AutoMigrate(model.Model)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Migrate successful")
}

func (server *Server) InitCommands(config AppConfig, dbConfig DBConfig) {
	server.initializeDB(dbConfig)

	cmdApp := cli.NewApp()
	cmdApp.Commands = []cli.Command{
		{
			Name: "db:migrate",
			Action: func(c *cli.Context) error {
				server.dbMigrate()
				return nil
			},
		},
		{
			Name: "db:seed",
			Action: func(c *cli.Context) error {
				err := seeder.DBSeed(server.DB)
				if err != nil {
					log.Fatal(err)
				}

				return nil
			},
		},
	}

	err := cmdApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
