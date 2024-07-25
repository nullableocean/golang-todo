package main

import (
	"github.com/joho/godotenv"
	"github.com/nullableocean/golang-todo"
	"github.com/nullableocean/golang-todo/internal/handler"
	"github.com/nullableocean/golang-todo/internal/repository"
	"github.com/nullableocean/golang-todo/internal/services"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error init godotenv: %s", err.Error())
	}

	if err := initConfig(); err != nil {
		logrus.Fatalf("error init config: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB_NAME"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("error from db: %s", err.Error())
	}

	repo := repository.NewRepository(db)
	service := services.NewServices(repo)
	handlers := handler.NewHandler(service)

	serv := new(todo.Server)
	err = serv.Run(viper.GetString("port"), handlers.InitRoutes())
	if err != nil {
		logrus.Fatalf("error from serv: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
