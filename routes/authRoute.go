package routes

import (
	"github.com/gin-gonic/gin"
	controllers "udit.com/blog/controllers"
	"udit.com/blog/middleware"
)

func AdminRoutes(rg *gin.Engine) {
	router := rg.Group("admin")
	router.POST("/signUp", controllers.SignInUser)
	router.Use(middleware.DeserializeUser())

	router.POST("/createBlog", controllers.CreateBlog)
	router.PUT("/blog/:blogId", controllers.EditBlog)
	router.DELETE("/blog/:blogId", controllers.DeleteBlog)
}
