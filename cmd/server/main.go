package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/marcosduarte-dev/TaskChrono-Back/configs"
	_ "github.com/marcosduarte-dev/TaskChrono-Back/docs"
	"github.com/marcosduarte-dev/TaskChrono-Back/internal/entity"
	dbProject "github.com/marcosduarte-dev/TaskChrono-Back/internal/infra/database/Project"
	dbTask "github.com/marcosduarte-dev/TaskChrono-Back/internal/infra/database/Task"
	dbTimer "github.com/marcosduarte-dev/TaskChrono-Back/internal/infra/database/Timer"
	"github.com/marcosduarte-dev/TaskChrono-Back/internal/infra/webserver/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// @title Go Task Chrono API
// @version 1.0
// @description BackEnd to timer project
// @termsOfService http://www.swagger.io/terms/

// @contact.name Marcos Duarte
// @contact.url http://github.com/marcosduarte-dev/
// @contact.email pe.marcos30@gmail.com

// @license.name MarkDev License
// @license.url http://github.com/marcosduarte-dev/

// @host localhost:8000
// @BasePath /
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
	db.AutoMigrate(&entity.Project{}, &entity.Task{}, &entity.Timer{})

	projectDB := dbProject.NewProject(db)
	ProjectHandler := handlers.NewProjectHandler(projectDB)

	taskDB := dbTask.NewTask(db)
	TaskHandler := handlers.NewTaskHandler(taskDB)

	timerDB := dbTimer.NewTimer(db)
	TimerHandler := handlers.NewTimerHandler(timerDB)

	r  := chi.NewRouter()
	r.Use(middleware.Logger)

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
	r.Route("/timers", func(r chi.Router) {
		r.Post("/", TimerHandler.CreateTimer)
		r.Options("/", TimerHandler.Options)
		r.Get("/", TimerHandler.GetTimers)
		r.Get("/{id}", TimerHandler.GetTimer)
		r.Get("/date/{date}", TimerHandler.GetTimerByDate)
		r.Get("/task/{task_id}", TimerHandler.GetTimerByTaskID)
		r.Put("/{id}", TimerHandler.UpdateTimer)
		r.Delete("/{id}", TimerHandler.DeleteTimer) 
		r.Options("/{id}", TimerHandler.Options) 
	})

	
	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	http.ListenAndServe(":8000", r)
}