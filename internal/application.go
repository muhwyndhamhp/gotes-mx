package internal

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Application struct {
	DB            *gorm.DB
	E             *echo.Echo
	RequireAuthMW echo.MiddlewareFunc
	CheckAuthMW   echo.MiddlewareFunc
}
