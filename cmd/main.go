package main

import (
	"github.com/joho/godotenv"
	"log/slog"
	"os"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"

	"github.com/Grishun/todo"
	"github.com/Grishun/todo/pkg/handler"
	"github.com/Grishun/todo/pkg/repository"
	"github.com/Grishun/todo/pkg/service"
)

func main() {

	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	if err := initConfig(); err != nil {
		log.Error("failed to init env config", err)
		os.Exit(1)
	}

	if err := godotenv.Load(); err != nil {
		log.Error("failed to read env files", err)
		os.Exit(1)
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		log.Error("failed to init db", err)
		os.Exit(1)
	}

	rep := repository.NewRep(db)
	srvs := service.NewService(rep)
	api := handler.NewHandler(srvs)

	srv := new(todo.Server)
	if err = srv.Run(viper.GetString("port"), api.InitRoutes()); err != nil {
		log.Error("failed to init db", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
