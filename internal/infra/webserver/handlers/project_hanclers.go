package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	dto "github.com/marcosduarte-dev/TaskChrono-Back/internal/dto/Project"
	"github.com/marcosduarte-dev/TaskChrono-Back/internal/entity"
	database "github.com/marcosduarte-dev/TaskChrono-Back/internal/infra/database/Project"
	pkgEntity "github.com/marcosduarte-dev/TaskChrono-Back/pkg/entity"
	"github.com/marcosduarte-dev/TaskChrono-Back/pkg/errors"
	cors "github.com/marcosduarte-dev/TaskChrono-Back/pkg/func"
)

type ProjectHandler struct {
	ProjectDB database.ProjectInterface
}
func NewProjectHandler(db database.ProjectInterface) *ProjectHandler {
	return &ProjectHandler{db}
}

func (h *ProjectHandler) Options(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	w.WriteHeader(http.StatusOK)
}

// Create       Project   godoc
// @Summary     Create project
// @Description Create projectss
// @Tags        projects
// @Accept      json
// @Produce     json
// @Param       request   body     dto.ProjectInputDTO true "project request"
// @Success     201       {object} entity.Return
// @Failure     500       {object} entity.Return
// @Failure     400       {object} entity.Return
// @Router      /projects [post]
func (h *ProjectHandler) CreateProject(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	var project dto.ProjectInputDTO
	err := json.NewDecoder(r.Body).Decode(&project)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		error := pkgEntity.Return{Status: "Error", Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	p, err := entity.NewProject(project.Name, project.Color, project.Description, project.UserID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		error := pkgEntity.Return{Status: "Error", Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	err = h.ProjectDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		error := pkgEntity.Return{Status: "Error", Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	ret := pkgEntity.Return{Status: "Success", Message: "Project Created"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ret)
}

func (h *ProjectHandler) GetProjects(w http.ResponseWriter, r *http.Request) {
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

	projects, err := h.ProjectDB.FindMyProjects(pageInt, limitInt, sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		error := pkgEntity.Return{Status: "Error", Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(projects)
}

func (h *ProjectHandler) GetProject(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		error := pkgEntity.Return{Status: "Error", Message: errors.ErrIDIsRequired.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	project, err := h.ProjectDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		error := pkgEntity.Return{Status: "Error", Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(project)
}

func (h *ProjectHandler) UpdateProject(w http.ResponseWriter, r *http.Request)  {
	cors.EnableCors(&w)
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		error := pkgEntity.Return{Status: "Error", Message: errors.ErrIDIsRequired.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	var project entity.Project
	err := json.NewDecoder(r.Body).Decode(&project)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		error := pkgEntity.Return{Status: "Error", Message: errors.ErrJSON.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	project.ID, err = pkgEntity.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		error := pkgEntity.Return{Status: "Error", Message: errors.ErrInvalidID.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	_, err = h.ProjectDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		error := pkgEntity.Return{Status: "Error", Message: errors.ErrProjectNotFound.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	err = h.ProjectDB.Update(&project)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		error := pkgEntity.Return{Status: "Error", Message: errors.ErrProjectUpdate.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	ret := pkgEntity.Return{Status: "Success", Message: "Project Updated"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ret)
}

func (h *ProjectHandler) DeleteProject(w http.ResponseWriter, r *http.Request) {
	cors.EnableCors(&w)
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		error := pkgEntity.Return{Status: "Error", Message: errors.ErrIDIsRequired.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	_, err := h.ProjectDB.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		error := pkgEntity.Return{Status: "Error", Message: errors.ErrProjectNotFound.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	err = h.ProjectDB.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		error := pkgEntity.Return{Status: "Error", Message: errors.ErrProjectDelete.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	ret := pkgEntity.Return{Status: "Success", Message: "Project Deleted"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ret)
}