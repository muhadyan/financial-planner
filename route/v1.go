package route

import (
	"github.com/labstack/echo"
	"github.com/muhadyan/financial-planner/middleware"
	"github.com/muhadyan/financial-planner/utils"
)

func V1Routes(g *echo.Group, controllers AppModels) {
	g.GET("/example", controllers.Example.GetExampleName)

	user := g.Group("/user")
	user.POST("/signup", controllers.User.SignUp)
	user.GET("/verify", controllers.User.Verify)
	user.POST("/login", controllers.User.LogIn)

	gold := g.Group("/gold", middleware.JWTVerify([]string{utils.RoleUser}))
	gold.POST("", controllers.Gold.Create)
}
