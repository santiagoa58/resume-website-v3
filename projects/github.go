package projects

import (
	"os"
)

type Projects struct {
	ProfileURL string
}

func GetProjects() *Projects {
	projects := Projects{
		ProfileURL: os.Getenv("PROJECTS_PROFILE_URL"),
	}
	return &projects
}
