package main

import (
	"database/sql"
	"io/ioutil"
	"log"
	"path/filepath"
	"pr-reviewer-service/internal/config"
	"pr-reviewer-service/internal/handler"
	"pr-reviewer-service/internal/repository"
	"pr-reviewer-service/internal/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Не удалось загрузить конфиг: %v", err)
	}

	db, err := sql.Open("postgres", cfg.GetDSN())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal("Не удалось подключиться к БД:", err)
	}
	log.Println("Successfully connected to PostgreSQL!")

	migrationPath := filepath.Join("migrations", "0001_initial_schema.sql")
	if data, err := ioutil.ReadFile(migrationPath); err != nil {
		log.Printf("Warning: не найден файл миграции: %v", err)
	} else {
		if _, err := db.Exec(string(data)); err != nil {
			log.Printf("Warning: ошибка при применении миграции: %v", err)
		} else {
			log.Println("Миграции успешно применены")
		}
	}

	repo := repository.NewRepository(db)
	userService := service.NewUserService(repo)
	prService := service.NewPRService(repo)

	userHandler := handler.NewUserHandler(userService)
	prHandler := handler.NewPRHandler(prService)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	v1 := e.Group("/api/v1")

	v1.POST("/users", userHandler.CreateUser)
	v1.GET("/users/:id", userHandler.GetUser)
	v1.PUT("/users/:id", userHandler.UpdateUser)
	v1.DELETE("/users/:id", userHandler.DeleteUser)
	v1.GET("/users", userHandler.ListUsers)

	v1.POST("/pull-requests", prHandler.CreatePR)

	log.Printf("Сервер запущен на :%s", cfg.HTTPPort)
	log.Fatal(e.Start(":" + cfg.HTTPPort))
}
