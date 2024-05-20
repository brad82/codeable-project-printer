package codeable

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"slices"
	"time"
)

var cache = []int{}

type ProjectClient struct {
	Token string
}

func (a *ProjectClient) getProjects() []Project {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", "https://api.codeable.io/experts/projects/new-projects?page=1", nil)
	req.Header.Add("Authorization", "Bearer "+a.Token)
	req.Header.Add("Accept", "application/json, application/vnd.codeable.v1")

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

func (a *ProjectClient) poll(ch chan []Project) {
	// Load, cache and discard the first results
	_ = a.getNewProjects()

	for {
		log.Print("Scanning for new projects")
		projects := a.getNewProjects()
		if len(projects) > 0 {
			ch <- projects
		}

		time.Sleep(5 * time.Minute)
	}
}

func (a *ProjectClient) StartPoll(interval int) chan []Project {
	ch := make(chan []Project)

	if interval <= 0 {
		panic("Interval must be greater than 0")
	}

	go a.poll(ch)

	return ch
}
