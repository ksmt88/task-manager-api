package controllers

import (
	"net/http"
	"strconv"

	"github.com/ksmt88/task-manager-api/internal/task/domain"
	"github.com/ksmt88/task-manager-api/internal/task/interfaces/database"
	"github.com/ksmt88/task-manager-api/internal/task/usecase"
	"github.com/labstack/echo"
)

type CategoryController struct {
	Interactor usecase.CategoryInteractor
}

func NewCategoryController(sqlHandler database.SqlHandler) *CategoryController {
	return &CategoryController{
		Interactor: usecase.CategoryInteractor{
			CategoryRepository: &database.CategoryRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *CategoryController) Create(c echo.Context) error {
	var category domain.Category
	err := c.Bind(&category)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: "failed to bind data.",
			Detail:  err,
		})
	}
	id, err := controller.Interactor.Add(category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: "failed to create category.",
			Detail:  err,
		})
	}
	category.Id = id
	return c.JSON(http.StatusCreated, domain.Response{ResponseData: category})
}

func (controller *CategoryController) Index(c echo.Context) error {
	var categories domain.Categories
	categories, err := controller.Interactor.All()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: "failed to get categories.",
			Detail:  err,
		})
	}
	return c.JSON(http.StatusOK, domain.Response{ResponseData: categories})
}

func (controller *CategoryController) Show(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	category, err := controller.Interactor.FindById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: "failed to get category.",
			Detail:  err,
		})
	}
	return c.JSON(http.StatusOK, domain.Response{ResponseData: category})
}

func (controller *CategoryController) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var category domain.Category
	err := c.Bind(&category)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: "failed to bind data.",
			Detail:  err,
		})
	}
	if err := controller.Interactor.Save(id, category); err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: "failed to update category.",
			Detail:  err,
		})
	}
	return c.JSON(http.StatusOK, domain.Response{ResponseData: "OK"})
}

func (controller *CategoryController) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := controller.Interactor.Remove(id); err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: "failed to delete category.",
			Detail:  err,
		})
	}
	return c.JSON(http.StatusOK, domain.Response{ResponseData: "OK"})
}
