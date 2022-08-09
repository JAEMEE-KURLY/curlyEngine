package httpserver

import (
	"strings"

	. "almcm.poscoict.com/scm/pme/curly-engine/log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	echoSwagger "github.com/swaggo/echo-swagger"
	"almcm.poscoict.com/scm/pme/curly-engine/docs"

	_ "almcm.poscoict.com/scm/pme/curly-engine/docs"
	"almcm.poscoict.com/scm/pme/curly-engine/configure"
)

// HttpServer
// @title PosGo REST API
// @version 0.1.0
// @BasePath /api/v1
// @query.collection.format multi
//
// @description <h2><b>PosGo REST API Swagger Documentation</b></h2>
//
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
//
// @tag.name Sample
// @tag.description Sample TAG
func HttpServer(port string) {
	conf := configure.GetConfig()

	port = strings.Trim(port, " ")
	if len(port) <= 0 {
		port = conf.CurlyEngine.HttpServerPort
	}
	docs.SwaggerInfo.Host = conf.CurlyEngine.SwaggerAddress + ":" + port

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Logger.SetOutput(GetLogWriter())

	// swagger documents
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	setRoute(e)

	err := e.Start(":" + port)
	if err != nil {
		e.Logger.Fatal(err)
	}
}

// setRoute
func setRoute(e *echo.Echo) {
	// Insert API Route
	// ApiUser.SetRoute(e)
}