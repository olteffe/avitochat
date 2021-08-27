package server

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/olteffe/avitochat/internal/handlers"
	"github.com/olteffe/avitochat/internal/repository"
	"github.com/olteffe/avitochat/internal/usecase"
	"gorm.io/driver/postgres"
)

type Config struct {
	Host       string `json:"host"`
	Port       int    `json:"port"`
	DBHost     string `json:"dbHost"`
	DBUser     string `json:"dbUser"`
	DBPassword string `json:"dbPassword"`
	DBName     string `json:"dbName"`
	DBPort     int    `json:"dbPort"`
}

// StartServer func
func StartServer(quit chan os.Signal, config Config) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		config.DBHost,
		config.DBUser,
		config.DBPassword,
		config.DBName,
		config.DBPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("cannot connect to databse: %v", err.Error())
	}
	postgresDB, err := db.DB()
	if err != nil {
		log.Fatalf("cannot connect to databse: %v", err.Error())
	}
	defer postgresDB.Close()
	err = postgresDB.Ping()
	if err != nil {
		log.Fatalf("cannot connect to databse: %v", err.Error())
	}

	repos := repository.NewRepository(db)
	usecases := usecase.NewService(repos)
	handler := handlers.NewHandler(usecases)
	e := echo.New()
	handler.Init(e)

	// Start server
	go func() {
		addr := fmt.Sprintf("%s:%d", config.Host, config.Port)
		if err := e.Start(addr); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
