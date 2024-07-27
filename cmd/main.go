package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/nullableocean/golang-todo"
	"github.com/nullableocean/golang-todo/internal/handler"
	"github.com/nullableocean/golang-todo/internal/repository"
	"github.com/nullableocean/golang-todo/internal/services"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

// @title Todo API
// @version 0.0.1
// @description Same learn Go project

// @BasePath /api
// @produce json

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name authorization
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
	go func() {
		err = serv.Run(viper.GetString("port"), handlers.InitRoutes())
		if err != nil {
			logrus.Fatalf("error from serv: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	logrus.Println("app shutting down...")

	err = serv.Shutdown(context.Background())
	if err != nil {
		logrus.Fatalf("server shutting down with error: %s", err.Error())
	}

	err = db.Close()
	if err != nil {
		logrus.Fatalf("database closing with error: %s", err.Error())
	}

	logrus.Println("shutting down success")
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
