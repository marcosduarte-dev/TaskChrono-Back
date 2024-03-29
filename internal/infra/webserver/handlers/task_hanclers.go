package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	dto "github.com/marcosduarte-dev/TaskChrono-Back/internal/dto/Task"
	"github.com/marcosduarte-dev/TaskChrono-Back/internal/entity"
	database "github.com/marcosduarte-dev/TaskChrono-Back/internal/infra/database/Task"
	pkgEntity "github.com/marcosduarte-dev/TaskChrono-Back/pkg/entity"
	"github.com/marcosduarte-dev/TaskChrono-Back/pkg/errors"
	cors "github.com/marcosduarte-dev/TaskChrono-Back/pkg/func"
)

type TaskHandler struct {
	TaskDB database.TaskInterface
}

func NewTaskHandler(db database.TaskInterface) *TaskHandler {
	return &TaskHandler{db}
}

func (h *TaskHandler) Options(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	w.WriteHeader(http.StatusOK)
}

// Create       Task   godoc
// @Summary     Create task
// @Description Create tasks
// @Tags        tasks
// @Accept      json
// @Produce     json
// @Param       request   body     dto.TaskInputDTO true "task request"
// @Success     201       {object} entity.Return
// @Failure     500       {object} entity.Return
// @Failure     400       {object} entity.Return
// @Router      /tasks [post]
func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	var task dto.TaskInputDTO
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		error := pkgEntity.Return{Status: "Error", Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	p, err := entity.NewTask(task.Name, task.Color, task.Description, task.ProjectID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		error := pkgEntity.Return{Status: "Error", Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	err = h.TaskDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		error := pkgEntity.Return{Status: "Error", Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	ret := pkgEntity.Return{Status: "Success", Message: "Task Created"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ret)
}

// ListTasks godoc
// @Summary     List tasks
// @Description get all tasks
// @Tags        tasks
// @Accept      json
// @Produce     json
// @Param       page      query    string         false "page number"
// @Param       limit     query    string         false "limit"
// @Success     200       {array}  entity.Task
// @Failure     404       {object} entity.Return
// @Failure     500       {object} entity.Return
// @Router      /tasks [get]
func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 0
	}

	sort := r.URL.Query().Get("sort")

	tasks, err := h.TaskDB.FindMyTasks(pageInt, limitInt, sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		error := pkgEntity.Return{Status: "Error", Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

// GetTask godoc
// @Summary     Get a task 
// @Description Get a task by id
// @Tags        tasks
// @Accept      json
// @Produce     json
// @Param 			id 						 path 		string 				 true "task ID" Format(uuid)
// @Success 		200 					 {object} entity.Task
// @Failure 		400 					 {object} entity.Return
// @Failure 		404 					 {object} entity.Return
// @Router      /tasks/{id} [get]
func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		error := pkgEntity.Return{Status: "Error", Message: errors.ErrIDIsRequired.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	task, err := h.TaskDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		error := pkgEntity.Return{Status: "Error", Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

// GetTaskByProjectID godoc
// @Summary 		Get tasks by project ID
// @Description Get tasks by project ID
// @Tags 				tasks
// @Accept 			json
// @Produce 		json
// @Param 			id 						 path 		string 								 true "Project ID" Format(uuid)
// @Success 		200						 {array}  entity.Task
// @Failure 		400 					 {object} entity.Return
// @Failure 		404 					 {object} entity.Return
// @Router 			/tasks/project/{id} [get]
func (h *TaskHandler) GetTaskByProjectID(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	projectID := chi.URLParam(r, "project_id")
	if projectID == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		error := pkgEntity.Return{Status: "Error", Message: errors.ErrIDIsRequired.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	task, err := h.TaskDB.FindByProjectID(projectID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		error := pkgEntity.Return{Status: "Error", Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

// Update Task godoc
// @Summary 		Update a task
// @Description Update a task by ID
// @Tags 				tasks
// @Accept 			json
// @Produce 		json
// @Param 			id 						 path 		string 								 true "task ID" Format(uuid)
// @Param 			request 			 body 		dto.TaskInputDTO true "task request"
// @Success 		200						 {object} entity.Return
// @Failure 		400 					 {object} entity.Return
// @Failure 		404 					 {object} entity.Return
// @Failure 		500 					 {object} entity.Return
// @Router 			/tasks/{id} [put]
func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request)  {
	cors.EnableCors(&w)
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		error := pkgEntity.Return{Status: "Error", Message: errors.ErrIDIsRequired.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	var task entity.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		error := pkgEntity.Return{Status: "Error", Message: errors.ErrJSON.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	task.ID, err = pkgEntity.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		error := pkgEntity.Return{Status: "Error", Message: errors.ErrInvalidID.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	_, err = h.TaskDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		error := pkgEntity.Return{Status: "Error", Message: errors.ErrTaskNotFound.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	err = h.TaskDB.Update(&task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		error := pkgEntity.Return{Status: "Error", Message: errors.ErrTaskUpdate.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	ret := pkgEntity.Return{Status: "Success", Message: "Task Updated"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ret)
}

// Delete Task godoc
// @Summary 		Delete a task
// @Description Delete a task by ID
// @Tags 				tasks
// @Accept 			json
// @Produce 		json
// @Param 			id 						 path 		string 								 true "task ID" Format(uuid)
// @Success 		200					   {object} entity.Return
// @Failure 		400 					 {object} entity.Return
// @Failure 		404 					 {object} entity.Return
// @Failure 		500 					 {object} entity.Return
// @Router 			/tasks/{id} [delete]
func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		error := pkgEntity.Return{Status: "Error", Message: errors.ErrIDIsRequired.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	_, err := h.TaskDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		error := pkgEntity.Return{Status: "Error", Message: errors.ErrTaskNotFound.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	err = h.TaskDB.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		error := pkgEntity.Return{Status: "Error", Message: errors.ErrTaskDelete.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	ret := pkgEntity.Return{Status: "Success", Message: "Task Deleted"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ret)
}