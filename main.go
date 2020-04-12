package main

import (
	"log"
	"os"

	"github.com/devrodriguez/audit-api/controllers"
	"github.com/devrodriguez/audit-api/db"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	server := gin.New()

	log.Println(port)
	// Middlewares

	// Db connection
	err := db.Connect()
	if err != nil {
		log.Fatal(err)
		return
	}

	apiRoutes := server.Group("/api")
	{
		// Auditor routes
		apiRoutes.GET("/auditors", controllers.GetAuditor)
		apiRoutes.POST("/auditors", controllers.AddAuditor)
		apiRoutes.PUT("/auditors/:id", controllers.UpdateAuditor)
		apiRoutes.DELETE("/auditors", controllers.DeleteAuditor)

		// Enterprise routes
		apiRoutes.GET("/enterprises", controllers.GetEnterprises)
		apiRoutes.POST("/enterprises", controllers.AddEnterprise)
		apiRoutes.PUT("/enterprises/:id", controllers.UpdateEnterprise)
	}

	// Start server
	if port == "" {
		port = "3001"
	}

	// Run server
	server.Run(":" + port)

}
