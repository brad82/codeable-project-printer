package main

import (
	codeable "github.com/brad82/codeable-project-printer/codeable"
	"github.com/brad82/codeable-project-printer/escpos"
)

func printProject(receipt *escpos.Printer, p codeable.Project) {
	receipt.Rule("*")
	receipt.Sprintln("ID:     %d", p.ID)
	receipt.Sprintln("Client: %s (%.1f)", p.Client.FullName, p.Client.AverageReviewRating)
	receipt.Sprintln("Counts: %d Comments | %d Estimates", p.PublicCommentsCount, p.EstimatesCount)
	receipt.Sprintln("Budget: %s", p.Budget)
	receipt.Rule("=")
	receipt.Sprintln(p.Title)
	receipt.Rule("=")
	receipt.Sprintln(p.PostedDate.String())
	receipt.Rule("*")
	receipt.Feed(10).Cut()

	receipt.Flush()
}
