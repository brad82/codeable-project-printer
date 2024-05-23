package codeable

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"slices"
	"time"
)

const CDBL_API_URL = "https://api.codeable.io"

var cache = []int{}

type ProjectClient struct {
	authToken string
}

func withJsonHeaders(req *http.Request) *http.Request {
	req.Header.Add("Accept", "application/json, application/vnd.codeable.v1")

	return req
}
func (a *ProjectClient) getProjects() []Project {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", CDBL_API_URL+"/experts/projects/new-projects?page=1", nil)

	withJsonHeaders(req)

	req.Header.Add("Authorization", "Bearer "+a.authToken)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Println(resp.StatusCode)
		return nil
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var projects []Project
	json.Unmarshal(bodyBytes, &projects)

	return projects
}

func (a *ProjectClient) getNewProjects() []Project {
	var projects []Project

	for _, project := range a.getProjects() {
		cached := slices.Contains(cache, project.ID)
		if cached {
			continue
		}

		cache = append(cache, project.ID)
		projects = append(projects, project)
	}

	return projects
}

func (a *ProjectClient) poll(ch chan Project, interval int) {
	// Load, cache and discard the first results
	_ = a.getNewProjects()

	for {
		log.Print("Scanning for new projects")
		projects := a.getNewProjects()

		if len(projects) > 0 {

			log.Printf("Found %d Projects", len(projects))
			for _, project := range projects {
				ch <- project
			}
		}

		time.Sleep(time.Duration(interval) * time.Minute)
	}
}

func (a *ProjectClient) StartPoll(interval int) chan Project {
	ch := make(chan Project, 24)

	if interval <= 0 {
		panic("Interval must be greater than 0")
	}

	go a.poll(ch, interval)

	return ch
}
