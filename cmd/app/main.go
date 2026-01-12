package main

import (
	"fmt"
	"log"
	"net/http"
	"resume-website/internal/projects"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Projects struct {
	ProfileURL string
}

func healthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "healthy",
		"status":  "ok",
	})
}

func getSocialLinks(c *gin.Context) {
	proj := projects.GetProjects()
	c.String(http.StatusOK, "%s", proj.ProfileURL)
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file")
	}
}

func initRouter() *gin.Engine {
	// equivalent to:
	// r := gin.New()
	// r.Use(gin.Logger())
	// r.Use(gin.Recovery())
	r := gin.Default()
	r.GET("/health", healthCheckHandler)
	r.GET("/socials", getSocialLinks)
	return r
}

func startServer(router *gin.Engine) {
	fmt.Println("Starting server on :8081")
	err := router.Run(":8081")
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}

func main() {
	router := initRouter()
	startServer(router)
}
