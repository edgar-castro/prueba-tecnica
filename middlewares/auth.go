package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("AUTH_KEY")
	user_id := session.Get(("USER_ID"))
	fmt.Println(user_id)
	if user == nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "You dont have permisosn to be here, try to login firtst",
		})
		c.Abort()
		return
	}

	c.Next()
}
