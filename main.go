package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/muhwyndhamhp/gotes-mx/config"
	"github.com/muhwyndhamhp/gotes-mx/db"
	"github.com/muhwyndhamhp/gotes-mx/internal"
	"github.com/muhwyndhamhp/gotes-mx/internal/api"
	"github.com/muhwyndhamhp/gotes-mx/utils/routing"
	"github.com/muhwyndhamhp/gotes-mx/utils/template"
)

func main() {
	e := bootstrapRouter()

	app := &internal.Application{
		DB: db.GetDB(),
		E:  e,
	}

	api.RegisterRoutes(app)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.Get(config.APP_PORT))))
}

func bootstrapRouter() *echo.Echo {
	e := echo.New()

	routing.SetupRouter(e)
	template.NewTemplateRenderer(e)

	e.Static("/public", "public")

	return e
}
