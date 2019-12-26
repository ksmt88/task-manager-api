package infrastructure

import (
	"net/http"

	"github.com/ksmt88/task-manager-api/internal/task/interfaces/controllers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var Router *echo.Echo

func init() {
	e := echo.New()

	/*accessLog, err := os.OpenFile("./log/access.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: accessLog,
	}))*/
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	/*appLog, err := os.OpenFile("./log/app.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(appLog)*/
	sqlHandler := NewHandler()
	projectController := controllers.NewProjectController(sqlHandler)
	e.GET("/projects", projectController.Index)
	e.GET("/project/:id", projectController.Show)
	e.POST("/project", projectController.Create)
	e.PATCH("/project/:id", projectController.Update)
	e.DELETE("/project/:id", projectController.Delete)

	categoryController := controllers.NewCategoryController(sqlHandler)
	e.GET("/categories", categoryController.Index)
	e.GET("/category/:id", categoryController.Show)
	e.POST("/category", categoryController.Create)
	e.PATCH("/category/:id", categoryController.Update)
	e.DELETE("/category/:id", categoryController.Delete)

	taskController := controllers.NewTaskController(sqlHandler)
	e.GET("/tasks", taskController.Index)
	e.GET("/task/:id", taskController.Show)
	e.POST("/task", taskController.Create)
	e.PATCH("/task/:id", taskController.Update)
	e.DELETE("/task/:id", taskController.Delete)

	Router = e
}
