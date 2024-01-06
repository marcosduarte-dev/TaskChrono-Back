package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/marcosduarte-dev/TaskChrono-Back/configs"
	"github.com/marcosduarte-dev/TaskChrono-Back/internal/entity"
	database "github.com/marcosduarte-dev/TaskChrono-Back/internal/infra/database/Project"
	"github.com/marcosduarte-dev/TaskChrono-Back/internal/infra/webserver/handlers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Project{})

	projectDB := database.NewProject(db)
	ProjectHandler := handlers.NewProjectHandler(projectDB)

	r  := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/projects/", ProjectHandler.CreateProject)
	r.Options("/projects/", ProjectHandler.Options)
	r.Get("/projects/", ProjectHandler.GetProjects)
	r.Get("/projects/{id}", ProjectHandler.GetProject)
	r.Put("/projects/{id}", ProjectHandler.UpdateProject)
	r.Delete("/projects/{id}", ProjectHandler.DeleteProject)

	http.ListenAndServe(":8000", r)
}