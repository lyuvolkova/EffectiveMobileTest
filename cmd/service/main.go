package main

import (
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"log/slog"
	"net/http"
	"os"

	"effectiveMobileTest/internal/handler"
	"effectiveMobileTest/internal/service"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(fmt.Errorf("init dotenv: %w", err))
	}

	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})))

	db, err := gorm.Open(postgres.Open(os.Getenv("DB_DSN")), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	svc := service.New(db)
	slog.Info("server started localhost:8080")
	err = http.ListenAndServe(":8080", handler.New(svc))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
