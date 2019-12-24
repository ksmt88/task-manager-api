package controllers

import (
	"net/http"
	"strconv"

	"github.com/ksmt88/taskManager-api/internal/task/domain"
	"github.com/ksmt88/taskManager-api/internal/task/interfaces/database"
	"github.com/ksmt88/taskManager-api/internal/task/usecase"
	"github.com/labstack/echo"
)

type ProjectController struct {
	Interactor usecase.ProjectInteractor
}

func NewProjectController(sqlHandler database.SqlHandler) *ProjectController {
	return &ProjectController{
		Interactor: usecase.ProjectInteractor{
			ProjectRepository: &database.ProjectRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *ProjectController) Create(c echo.Context) error {
	var project domain.Project
	c.Bind(&project)
	id, err := controller.Interactor.Add(project)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: "failed to create project.",
			Detail:  err,
		})
	}
	project.Id = id
	return c.JSON(http.StatusCreated, domain.Response{ResponseData: project})
}

func (controller *ProjectController) Index(c echo.Context) error {
	var projects domain.Projects
	projects, err := controller.Interactor.All()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: "failed to get projects.",
			Detail:  err,
		})
	}
	return c.JSON(http.StatusOK, domain.Response{ResponseData: projects})
}

func (controller *ProjectController) Show(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	project, err := controller.Interactor.FindById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: "failed to get project.",
			Detail:  err,
		})
	}
	return c.JSON(http.StatusOK, domain.Response{ResponseData: project})
}

func (controller *ProjectController) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var project domain.Project
	c.Bind(&project)
	if err := controller.Interactor.Save(id, project); err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: "failed to update project.",
			Detail:  err,
		})
	}
	return c.JSON(http.StatusOK, domain.Response{ResponseData: "OK"})
}

func (controller *ProjectController) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := controller.Interactor.Remove(id); err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: "failed to delete project.",
			Detail:  err,
		})
	}
	return c.JSON(http.StatusOK, domain.Response{ResponseData: "OK"})
}
