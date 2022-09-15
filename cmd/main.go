package main

import (
	"log"

	"github.com/spf13/viper"
	restapi "github.com/viciousvs/restAPI"
	"github.com/viciousvs/restAPI/pkg/handler"
	"github.com/viciousvs/restAPI/pkg/repository"
	"github.com/viciousvs/restAPI/pkg/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("not read config %s", err.Error())
	}

	db, err := repository.NewPostgresDB(&repository.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "root",
		Password: "r4sh1td1n",
		DBName:   "postgres",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalf("not connect to Db: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	srv := new(restapi.Server)
	err = srv.Run(viper.GetString("port"), handler.InitRoutes())
	if err != nil {
		log.Fatalf("errir accured while running http server %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
