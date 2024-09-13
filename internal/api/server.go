package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/muhwyndhamhp/gotes-mx/internal"
	"github.com/muhwyndhamhp/gotes-mx/internal/src"
	"github.com/muhwyndhamhp/gotes-mx/utils/template"
)

func RegisterRoutes(app *internal.Application) {
	app.E.GET("/", func(c echo.Context) error {
		component := src.Index("Gotes-MX")
		return template.AssertRender(c, http.StatusOK, component)
	})

	app.E.GET("/click-me", func(c echo.Context) error {
		component := src.ClickMeBody()
		return template.AssertRender(c, http.StatusOK, component)
	})

	app.E.GET("/todos", func(c echo.Context) error {
		ctx := c.Request().Context()

		var todos []internal.Todo
		if err := app.DB.WithContext(ctx).Order("id ASC").Find(&todos).Error; err != nil {
			c.Logger().Error(err)
		}

		component := src.Todos(todos)
		return template.AssertRender(c, http.StatusOK, component)
	}, app.RequireAuthMW)

	app.E.POST("/todos/:id/toggle-finished", func(c echo.Context) error {
		ctx := c.Request().Context()
		id, _ := strconv.Atoi(c.Param("id"))

		var todo internal.Todo
		if err := app.DB.WithContext(ctx).Where("id = ?", id).First(&todo).Error; err != nil {
			c.Logger().Error(err)
		}

		if todo.IsFinished {
			todo.FinishedAt = nil
		} else {
			now := time.Now()
			todo.FinishedAt = &now
		}

		todo.IsFinished = !todo.IsFinished

		if err := app.DB.WithContext(ctx).Save(&todo).Error; err != nil {
			c.Logger().Error(err)
		}

		component := src.TodoRow(todo)
		return template.AssertRender(c, http.StatusOK, component)
	}, app.RequireAuthMW)
}
