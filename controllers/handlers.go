package controllers

import (
	// "go-test-api/middlewares"
	"fmt"
	"go-test-api/models"
	"go-test-api/utils"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LoginGETHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// Function that handle the login logic
func LoginPOSTHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		var resp models.SingInInput
		c.Bind(&resp)
		userInfo := models.GetUserByEmail(resp.Email)
		fmt.Println(userInfo)
		originalPassword := userInfo.Password
		if originalPassword == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Email have not registered",
			})
			return
		}
		passCoincide := utils.VerifyPassword(originalPassword, resp.Password)
		if passCoincide != nil {

			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Wrong pass",
			})
			return
		}
		session.Set("AUTH_KEY", true)
		session.Set("USER_ID", userInfo.ID)
		if err := session.Save(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Fail saving session",
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Loggin successfuly",
		})
		return
	}
}

// Function that handle the register loginc
func RegisterPOSTHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var resp models.SingUpInput
		c.Bind(&resp)
		emailExist := models.ExistEmail(resp.Email)
		if emailExist {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "The email is already registered",
			})
			return
		}
		hashedPass, err := utils.HashPassword(resp.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error in the registration",
			})
			return
		}
		user := models.User{
			Name:     resp.Name,
			Email:    resp.Email,
			Password: hashedPass,
		}
		err = models.CreateUser(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error while register the user",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "User register successfully",
		})

	}
}

// Function that handle the logout login
func LogOutHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		u := session.Get("AUTH_KEY")
		if u == nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "Logged Out (no session)",
			})
			return
		}
		session.Set("AUTH_KEY", "")
		session.Set("USER_ID", "")
		session.Clear()
		session.Options(sessions.Options{MaxAge: -1, Path: "/"})
		if err := session.Save(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Fail saving session",
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Logged Out",
		})
		return
	}
}

// Function that obtains all projects
func DashboardHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		key := session.Get("USER_ID")

		var projects []models.Project
		models.GetProjects(&projects, key.(int))
		c.JSON(http.StatusOK, projects)
	}
}

// Function that obtains an specific project by the id
func ProjectGETHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var resp models.Project
		if err := models.GetProject(&resp, id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error trying get project",
			})
			return
		}

		var project models.ProjectInput
		project.Name = resp.Name
		project.Description = resp.Description
		project.StartDate = resp.StartDate
		project.EndDate = resp.EndDate
		project.Budget = resp.Budget

		c.JSON(http.StatusOK, project)
	}
}

// Function that create a new project
func ProjectPOSTHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		var resp models.ProjectInput
		c.BindJSON(&resp)
		userId := session.Get("USER_ID")
		project := models.Project{
			Name:        resp.Name,
			Description: resp.Description,
			StartDate:   resp.StartDate,
			EndDate:     resp.EndDate,
			Budget:      resp.Budget,
			UserId:      userId.(int),
		}
		if err := models.CreateProject(&project); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error creating the project",
			})
			fmt.Print(err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Project created successfully",
		})
	}
}

// Function that update a project by the id
func ProjectPUTHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var resp models.ProjectInput
		c.Bind(&resp)
		if err := models.UpdateProject(&resp, id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"mesaage": "Error updating the project data",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Project updated successfuly",
		})
	}
}

func ProjectDELETEHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		if err := models.DeleteProject(id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"mesaage": "Error deleting the project data",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Project deleted successfuly",
		})
	}
}
