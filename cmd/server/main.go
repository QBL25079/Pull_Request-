package main

import (
	"database/sql"
	"log"
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

	repo := repository.NewUserRepository(db)
	userService := service.NewUserService(repo)
	h := handler.NewUserHandler(userService)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	v1 := e.Group("/api/v1")
	v1.POST("/users", h.CreateUser)
	v1.POST("/users", h.CreateUser)
	v1.GET("/users/:id", h.GetUser)
	v1.PUT("/users/:id", h.UpdateUser)
	v1.DELETE("/users/:id", h.DeleteUser)
	v1.GET("/users", h.ListUsers)

	//v1.POST("/teams", h.CreateTeam)
	//v1.PUT("/teams/:teamName", h.UpdateTeamName)
	//v1.GET("/users/:userID", h.GetUserByID)
	//v1.POST("/pull-requests", h.CreatePullRequest)
	//v1.GET("/pull-requests/:prID", h.GetPullRequestByID)

	log.Printf("Сервер запущен на :%s", cfg.HTTPPort)
	log.Fatal(e.Start(":" + cfg.HTTPPort))
}
