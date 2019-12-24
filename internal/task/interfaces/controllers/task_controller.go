package controllers

import (
	"net/http"
	"strconv"

	"../../domain"
	"../../interfaces/database"
	"../../usecase"
	"github.com/labstack/echo"
)

type TaskController struct {
	Interactor usecase.TaskInteractor
}

func NewTaskController(sqlHandler database.SqlHandler) *TaskController {
	return &TaskController{
		Interactor: usecase.TaskInteractor{
			TaskRepository: &database.TaskRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *TaskController) Create(c echo.Context) error {
	var task domain.Task
	c.Bind(&task)
	t, err := controller.Interactor.Add(task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: "failed to create task.",
			Detail:  err,
		})
	}
	return c.JSON(http.StatusCreated, domain.Response{ResponseData: t})
}

func (controller *TaskController) Index(c echo.Context) error {
	var tasks domain.Tasks
	tasks, err := controller.Interactor.All()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: "failed to get tasks.",
			Detail:  err,
		})
	}
	return c.JSON(http.StatusOK, domain.Response{ResponseData: tasks})
}

func (controller *TaskController) Show(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	task, err := controller.Interactor.FindById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: "failed to get task.",
			Detail:  err,
		})
	}
	return c.JSON(http.StatusOK, domain.Response{ResponseData: task})
}

func (controller *TaskController) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var task domain.Task
	c.Bind(&task)
	if err := controller.Interactor.Save(id, task); err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: "failed to update task.",
			Detail:  err,
		})
	}
	return c.JSON(http.StatusOK, domain.Response{ResponseData: "OK"})
}

func (controller *TaskController) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := controller.Interactor.Remove(id); err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: "failed to delete task.",
			Detail:  err,
		})
	}
	return c.JSON(http.StatusOK, domain.Response{ResponseData: "OK"})
}
