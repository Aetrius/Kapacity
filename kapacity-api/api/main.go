package api

import (
	Postgres "kapacity-api/internal/postgres"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	cors "github.com/rs/cors/wrapper/gin"
)

func StartServer() {
	r := gin.Default()

	godotenv.Load(".env")
	EnvironmentID := os.Getenv("ENVIRONMENT_ID")

	r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8081", "http://localhost:8081", "localhost:8081", "localhost:8080", "*"},
		AllowedHeaders:   []string{"Origin", "Content-Type", "Accept", "Host"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
		// Enable Debugging for testing, consider disabling in production
		Debug: false,
	}))

	dbConnection := Postgres.Connect()

	v1 := r.Group("/v1")

	setupNodeRoutes(v1.Group("/node"), dbConnection, EnvironmentID)
	setupClusterRoutes(v1.Group("/cluster"), dbConnection, EnvironmentID)
	setupApplicationRoutes(v1.Group("/application"), dbConnection, EnvironmentID)

	r.Run(":8080")
}
