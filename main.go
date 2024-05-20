package main

import (
	"fmt"
	"log"
	"os"

	codeable "github.com/brad82/codeable-project-printer/codeable"
)

func printProject(p codeable.Project) {
	receipt := Receipt{
		Width: 40,
	}

	receipt.Divider("*")
	receipt.Line("ID:     %d", p.ID)
	receipt.Line("Client: %s (%.1f)", p.Client.FullName, p.Client.AverageReviewRating)
	receipt.Line("Counts: %d Comments | %d Estimates", p.PublicCommentsCount, p.EstimatesCount)
	receipt.Line("Budget: %s", p.Budget)
	receipt.Rule()
	receipt.Line(p.Title)
	receipt.Rule()
	receipt.Line(p.PostedDate.String())
	receipt.Divider("*")

	output := receipt.Flush()
	fmt.Print(output + "\n\n")
}

func main() {

	if len(os.Args) != 2 {
		fmt.Printf("Usage %s api_token", os.Args[0])
		return
	}

	api := &codeable.ProjectClient{
		Token: os.Args[1],
	}

	for {
		projects := api.StartPoll(30)
		log.Printf("Received %d new projects", len(<-projects))
		for _, project := range <-projects {
			printProject(project)
		}
	}
}
