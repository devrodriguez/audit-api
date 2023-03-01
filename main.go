package main

import (
	"fmt"
	"github.com/devrodriguez/audit-api/db"
	"io"
	"log"
	"os"

	"github.com/devrodriguez/audit-api/controllers"
	"github.com/devrodriguez/audit-api/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load .env
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// Set environment variables
	os.Setenv("JWT_SECRET", "dev1986")
	os.Setenv("EXPIRATION", "30m")

	port := os.Getenv("PORT")
	server := gin.New()

	// Set log file
	setupLogOutput()

	// Middlewares
	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.CORSAllowed())

	// Db connection
	connErr := db.Connect()
	if connErr != nil {
		log.Fatal(connErr)
		return
	}

	fmt.Printf("[port:%s]", port)

	// Routes
	pubRoutes := server.Group("/api")
	{
		pubRoutes.GET("/signin", controllers.SignIn)
	}

	secRoutes := server.Group("/api/auth")
	secRoutes.Use(middlewares.ValidateAuth())
	{
		// Auditor routes
		secRoutes.GET("/auditors", controllers.GetAuditor)
		secRoutes.POST("/auditors", controllers.AddAuditor)
		secRoutes.PUT("/auditors/:id", controllers.UpdateAuditor)
		secRoutes.DELETE("/auditors", controllers.DeleteAuditor)

		// Enterprise routes
		secRoutes.GET("/enterprises", controllers.GetEnterprises)
		secRoutes.POST("/enterprises", controllers.AddEnterprise)
		secRoutes.PUT("/enterprises/:id", controllers.UpdateEnterprise)

		// Audit routes
		secRoutes.GET("/audits", controllers.GetAudit)
		secRoutes.POST("/audits", controllers.CreateAudit)
		secRoutes.PUT("/audits/:id/goal", controllers.AddGoal)
		secRoutes.DELETE("/audits/:auditId/goal/:goalId", controllers.RemoveGoal)
	}

	if port == "" {
		port = "3001"
	}

	// Run server
	server.Run(":" + port)

}

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
