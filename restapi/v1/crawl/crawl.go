package crawl

import "github.com/labstack/echo/v4"
import . "almcm.poscoict.com/scm/pme/curly-engine/restapi/v1"

func SetRoute(e *echo.Echo) {
    e.GET(ApiPath+"/info", getInfoHandler)
}
