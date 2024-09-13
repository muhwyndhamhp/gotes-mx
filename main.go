package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/zitadel/zitadel-go/v3/pkg/authentication"
	"github.com/zitadel/zitadel-go/v3/pkg/zitadel"

	"github.com/labstack/echo/v4"
	"github.com/muhwyndhamhp/gotes-mx/config"
	"github.com/muhwyndhamhp/gotes-mx/db"
	"github.com/muhwyndhamhp/gotes-mx/internal"
	"github.com/muhwyndhamhp/gotes-mx/internal/api"
	"github.com/muhwyndhamhp/gotes-mx/utils/errs"
	"github.com/muhwyndhamhp/gotes-mx/utils/routing"
	"github.com/muhwyndhamhp/gotes-mx/utils/template"

	openid "github.com/zitadel/zitadel-go/v3/pkg/authentication/oidc"
)

var (
	domain      = flag.String("domain", "localhost", "your ZITADEL instance domain (in the form: https://<instance>.zitadel.cloud or https://<yourdomain>)")
	key         = flag.String("key", "MasterkeyNeedsToHave32Characters", "encryption key")
	clientID    = flag.String("clientID", config.Get(config.ZITADEL_CLIENT_ID), "clientID provided by ZITADEL")
	redirectURI = flag.String("redirectURI", "http://localhost:4040/auth/callback", "redirectURI registered at ZITADEL")
)

func main() {
	e, rmw, cmw := bootstrapRouter()

	app := &internal.Application{
		DB:            db.GetDB(),
		E:             e,
		RequireAuthMW: rmw,
		CheckAuthMW:   cmw,
	}

	api.RegisterRoutes(app)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.Get(config.APP_PORT))))
}

func bootstrapRouter() (*echo.Echo, echo.MiddlewareFunc, echo.MiddlewareFunc) {
	e := echo.New()

	ctx := context.Background()

	authClient, err := authentication.New(ctx, zitadel.New(*domain, zitadel.WithInsecure("5756")), *key,
		openid.DefaultAuthentication(*clientID, *redirectURI, *key),
	)
	if err != nil {
		panic(errs.Wrap(err, "zitadel sdk could not initialize"))
	}

	mw := authentication.Middleware(authClient)

	rmw := echo.WrapMiddleware(mw.RequireAuthentication())
	cmw := echo.WrapMiddleware(mw.CheckAuthentication())

	routing.SetupRouter(e)
	template.NewTemplateRenderer(e)

	e.Static("/public", "public")
	e.Group("/auth").Any("/*", echo.WrapHandler(authClient))

	return e, rmw, cmw
}
