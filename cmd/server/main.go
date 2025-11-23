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
		log.Fatalf("Error loading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.GetDSN())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	log.Println("Successfully connected to PostgreSQL!")

	repo := repository.NewRepository(db)

	userService := service.NewUserService(repo)
	teamService := service.NewTeamService(repo)
	prService := service.NewPRService(repo)

	userHandler := handler.NewUserHandler(userService)
	teamHandler := handler.NewTeamHandler(teamService)
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

	v1.POST("/team", teamHandler.CreateTeam)
	v1.GET("/team", teamHandler.ListTeams)

	v1.POST("/pull-requests", prHandler.CreatePR)

	log.Printf("Server started on port %s", cfg.HTTPPort)
	log.Fatal(e.Start(":" + cfg.HTTPPort))
}
