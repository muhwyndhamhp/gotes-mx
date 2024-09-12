package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/muhwyndhamhp/gotes-mx/config"
	"github.com/muhwyndhamhp/gotes-mx/src"
	"github.com/muhwyndhamhp/gotes-mx/template"
	"github.com/muhwyndhamhp/gotes-mx/utils/routing"
)

func main() {
	e := echo.New()
	routing.SetupRouter(e)

	e.Static("/public", "public")

	template.NewTemplateRenderer(e)

	e.GET("/", func(c echo.Context) error {
		component := src.Index("Gotes-MX")
		return template.AssertRender(c, http.StatusOK, component)
	})

	e.GET("/click-me", func(c echo.Context) error {
		component := src.ClickMeBody()
		return template.AssertRender(c, http.StatusOK, component)
	})

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.Get(config.APP_PORT))))
}
