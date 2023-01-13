package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	controllers "udit.com/blog/controllers"
)

func Routes(rg *gin.Engine) {
	router := rg.Group("api")
	router.GET("/", welcome)
	router.GET("/allBlogs", controllers.GetAllBlogs)
	router.GET("/getById/:blogId", controllers.GetSingleBlog)

}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To API",
	})
	return
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": "Route Not Found",
	})
	return
}
