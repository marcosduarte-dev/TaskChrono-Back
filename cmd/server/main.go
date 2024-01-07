package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/marcosduarte-dev/TaskChrono-Back/configs"
	"github.com/marcosduarte-dev/TaskChrono-Back/internal/entity"
	dbProject "github.com/marcosduarte-dev/TaskChrono-Back/internal/infra/database/Project"
	dbTask "github.com/marcosduarte-dev/TaskChrono-Back/internal/infra/database/Task"
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
	db.AutoMigrate(&entity.Project{}, &entity.Task{})

	projectDB := dbProject.NewProject(db)
	ProjectHandler := handlers.NewProjectHandler(projectDB)

	taskDB := dbTask.NewTask(db)
	TaskHandler := handlers.NewTaskHandler(taskDB)

	r  := chi.NewRouter()
	r.Use(middleware.Logger)

	// r.Post("/projects/", ProjectHandler.CreateProject)
	// r.Options("/projects/", ProjectHandler.Options)
	// r.Get("/projects/", ProjectHandler.GetProjects)
	// r.Get("/projects/{id}", ProjectHandler.GetProject)
	// r.Put("/projects/{id}", ProjectHandler.UpdateProject)
	// r.Delete("/projects/{id}", ProjectHandler.DeleteProject) 
	// r.Options("/projects/{id}", ProjectHandler.Options) 

	r.Route("/projects", func(r chi.Router) {
		r.Post("/", ProjectHandler.CreateProject)
		r.Options("/", ProjectHandler.Options)
		r.Get("/", ProjectHandler.GetProjects)
		r.Get("/{id}", ProjectHandler.GetProject)
		r.Put("/{id}", ProjectHandler.UpdateProject)
		r.Delete("/{id}", ProjectHandler.DeleteProject) 
		r.Options("/{id}", ProjectHandler.Options) 
	})
	r.Route("/tasks", func(r chi.Router) {
		r.Post("/", TaskHandler.CreateTask)
		r.Options("/", TaskHandler.Options)
		r.Get("/", TaskHandler.GetTasks)
		r.Get("/{id}", TaskHandler.GetTask)
		r.Get("/project/{project_id}", TaskHandler.GetTaskByProjectID)
		r.Put("/{id}", TaskHandler.UpdateTask)
		r.Delete("/{id}", TaskHandler.DeleteTask) 
		r.Options("/{id}", TaskHandler.Options) 
	})

	http.ListenAndServe(":8000", r)
}