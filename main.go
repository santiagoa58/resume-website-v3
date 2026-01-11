package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"resume-website/projects"

	"github.com/joho/godotenv"
)

type Projects struct {
	ProfileURL string
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"status":"ok","message":"healthy"}`+"\n")
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env file")
	}
}

func init() {
	http.HandleFunc("/health", healthCheckHandler)
	fmt.Println("Starting server on :8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}

func main() {
	projects := projects.GetProjects()
	fmt.Println("GitHub Profile URL:", projects.ProfileURL)
}
