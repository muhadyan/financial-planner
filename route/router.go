package route

import (
	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	app := App()

	e := echo.New()
	V1Routes(e.Group("/v1"), app)

	return e
}
