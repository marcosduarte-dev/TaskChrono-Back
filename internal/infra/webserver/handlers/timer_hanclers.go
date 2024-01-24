package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	dto "github.com/marcosduarte-dev/TaskChrono-Back/internal/dto/Timer"
	"github.com/marcosduarte-dev/TaskChrono-Back/internal/entity"
	database "github.com/marcosduarte-dev/TaskChrono-Back/internal/infra/database/Timer"
	pkgEntity "github.com/marcosduarte-dev/TaskChrono-Back/pkg/entity"
	"github.com/marcosduarte-dev/TaskChrono-Back/pkg/errors"
	cors "github.com/marcosduarte-dev/TaskChrono-Back/pkg/func"
)

type TimerHandler struct {
	TimerDB database.TimerInterface
}

func NewTimerHandler(db database.TimerInterface) *TimerHandler {
	return &TimerHandler{db}
}

func (h *TimerHandler) Options(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	w.WriteHeader(http.StatusOK)
}

// Create       Timer   godoc
// @Summary     Create timer
// @Description Create timers
// @Tags        timers
// @Accept      json
// @Produce     json
// @Param       request   body     dto.TimerInputDTO true "timer request"
// @Success     201       {object} entity.Return
// @Failure     500       {object} entity.Return
// @Failure     400       {object} entity.Return
// @Router      /timers [post]
func (h *TimerHandler) CreateTimer(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	var timer dto.TimerInputDTO
	err := json.NewDecoder(r.Body).Decode(&timer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		error := pkgEntity.Return{Status: "Error", Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	p, err := entity.NewTimer(timer.StartTime, timer.EndTime, timer.TotalDuration, timer.RecordType, timer.TaskID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		error := pkgEntity.Return{Status: "Error", Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	err = h.TimerDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		error := pkgEntity.Return{Status: "Error", Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	ret := pkgEntity.Return{Status: "Success", Message: "Timer Created"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ret)
}

// ListTimers godoc
// @Summary     List timers
// @Description get all timers
// @Tags        timers
// @Accept      json
// @Produce     json
// @Param       page      query    string         false "page number"
// @Param       limit     query    string         false "limit"
// @Success     200       {array}  entity.Timer
// @Failure     404       {object} entity.Return
// @Failure     500       {object} entity.Return
// @Router      /timers [get]
func (h *TimerHandler) GetTimers(w http.ResponseWriter, r *http.Request) {
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

	timers, err := h.TimerDB.FindMyTimers(pageInt, limitInt, sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		error := pkgEntity.Return{Status: "Error", Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(timers)
}

// GetTimer godoc
// @Summary     Get a timer 
// @Description Get a timer by id
// @Tags        timers
// @Accept      json
// @Produce     json
// @Param 			id 						 path 		string 				 true "timer ID" Format(uuid)
// @Success 		200 					 {object} entity.Timer
// @Failure 		400 					 {object} entity.Return
// @Failure 		404 					 {object} entity.Return
// @Router      /timers/{id} [get]
func (h *TimerHandler) GetTimer(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		error := pkgEntity.Return{Status: "Error", Message: errors.ErrIDIsRequired.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	timer, err := h.TimerDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		error := pkgEntity.Return{Status: "Error", Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(timer)
}

// GetTimerByTaskID godoc
// @Summary 		Get timer by task ID
// @Description Get timer by task ID
// @Tags 				timers
// @Accept 			json
// @Produce 		json
// @Param 			id 						 path 		string 								 true "Task ID" Format(uuid)
// @Success 		200						 {array}  entity.Timer
// @Failure 		400 					 {object} entity.Return
// @Failure 		404 					 {object} entity.Return
// @Router 			/timers/task/{id} [get]
func (h *TimerHandler) GetTimerByTaskID(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	taskID := chi.URLParam(r, "task_id")
	if taskID == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		error := pkgEntity.Return{Status: "Error", Message: errors.ErrIDIsRequired.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	timer, err := h.TimerDB.FindByTaskID(taskID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		error := pkgEntity.Return{Status: "Error", Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(timer)
}

// GetTimerByDate godoc
// @Summary 		Get timer by date
// @Description Get timer by date
// @Tags 				timers
// @Accept 			json
// @Produce 		json
// @Param 			date 					 path 		string        true "Date"
// @Success 		200						 {array}  entity.Timer
// @Failure 		400 					 {object} entity.Return
// @Failure 		404 					 {object} entity.Return
// @Router 			/timers/date/{date} [get]
func (h *TimerHandler) GetTimerByDate(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	dateStr := chi.URLParam(r, "date")
	if dateStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		error := pkgEntity.Return{Status: "Error", Message: errors.ErrDateIsRequired.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		error := pkgEntity.Return{Status: "Error", Message: errors.ErrDateInvalidFormat.Error()}
		json.NewEncoder(w).Encode(error)
	}
	timer, err := h.TimerDB.FindByDate(date)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		error := pkgEntity.Return{Status: "Error", Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(timer)
}

// Update Timer godoc
// @Summary 		Update a timer
// @Description Update a timer by ID
// @Tags 				timers
// @Accept 			json
// @Produce 		json
// @Param 			id 						 path 		string 					  true "timer ID" Format(uuid)
// @Param 			request 			 body 		dto.TimerInputDTO true "timer request"
// @Success 		200						 {object} entity.Return
// @Failure 		400 					 {object} entity.Return
// @Failure 		404 					 {object} entity.Return
// @Failure 		500 					 {object} entity.Return
// @Router 			/timers/{id} [put]
func (h *TimerHandler) UpdateTimer(w http.ResponseWriter, r *http.Request)  {
	cors.EnableCors(&w)
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		error := pkgEntity.Return{Status: "Error", Message: errors.ErrIDIsRequired.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	var timer entity.Timer
	err := json.NewDecoder(r.Body).Decode(&timer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		error := pkgEntity.Return{Status: "Error", Message: errors.ErrJSON.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	timer.ID, err = pkgEntity.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		error := pkgEntity.Return{Status: "Error", Message: errors.ErrInvalidID.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	_, err = h.TimerDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		error := pkgEntity.Return{Status: "Error", Message: errors.ErrTimerNotFound.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	err = h.TimerDB.Update(&timer)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		error := pkgEntity.Return{Status: "Error", Message: errors.ErrTimerUpdate.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	ret := pkgEntity.Return{Status: "Success", Message: "Timer Updated"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ret)
}

// Delete Timer godoc
// @Summary 		Delete a timer
// @Description Delete a timer by ID
// @Tags 				timers
// @Accept 			json
// @Produce 		json
// @Param 			id 						 path 		string 								 true "timer ID" Format(uuid)
// @Success 		200					   {object} entity.Return
// @Failure 		400 					 {object} entity.Return
// @Failure 		404 					 {object} entity.Return
// @Failure 		500 					 {object} entity.Return
// @Router 			/timers/{id} [delete]
func (h *TimerHandler) DeleteTimer(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		error := pkgEntity.Return{Status: "Error", Message: errors.ErrIDIsRequired.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	_, err := h.TimerDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		error := pkgEntity.Return{Status: "Error", Message: errors.ErrTimerNotFound.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	err = h.TimerDB.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		error := pkgEntity.Return{Status: "Error", Message: errors.ErrTimerDelete.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	ret := pkgEntity.Return{Status: "Success", Message: "Timer Deleted"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ret)
}