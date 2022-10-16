package routes

import (
	"go-test-api/controllers"

	"github.com/gin-gonic/gin"
)

func PublicRoutes(g *gin.RouterGroup) {
	g.GET("/login", controllers.LoginGETHandler())
	g.POST("/login", controllers.LoginPOSTHandler())
	g.POST("/register", controllers.RegisterPOSTHandler())
	g.POST("/logout", controllers.LogOutHandler())
}

func PrivateRoutes(g *gin.RouterGroup) {
	g.GET("/dashboard", controllers.DashboardHandler())
	g.GET("/project/:id", controllers.ProjectGETHandler())
	g.POST("/project/", controllers.ProjectPOSTHandler())
	g.PUT("/project/:id", controllers.ProjectPUTHandler())
	g.DELETE("/project/:id", controllers.ProjectDELETEHandler())
}
