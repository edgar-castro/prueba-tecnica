package routes

import (
	"go-test-api/middlewares"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.Default()

	router.Static("/assets", "./assets")

	router.Use(middlewares.CORSMiddleware())
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
	})
	router.Use(sessions.Sessions("mysession", store))

	public := router.Group("/auth")
	PublicRoutes(public)

	private := router.Group("/user", middlewares.Auth)
	PrivateRoutes(private)

	return router
}
