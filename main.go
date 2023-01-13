package main

import (
	"log"

	"github.com/gin-gonic/gin"

	config "udit.com/blog/config"
	"udit.com/blog/controllers"
	routes "udit.com/blog/routes"
)

func main() {
	// Connect DB
	db := config.ConnectDB()

	// Init Router
	router := gin.Default()

	// Route Handlers / Endpoints
	routes.Routes(router)
	routes.AdminRoutes(router)
	controllers.InitiateDB(db)
	log.Fatal(router.Run(":3000"))
}
