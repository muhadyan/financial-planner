package route

import "github.com/labstack/echo"

func V1Routes(g *echo.Group, controllers AppModels) {
	g.GET("/example", controllers.Example.GetExampleName)

	user := g.Group("/user")
	user.POST("/signup", controllers.User.SignUp)
	user.GET("/verify", controllers.User.Verify)
}
