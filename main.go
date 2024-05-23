package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	codeable "github.com/brad82/codeable-project-printer/codeable"
	"github.com/brad82/codeable-project-printer/escpos"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	printer, err := escpos.NewUSB()
	if err != nil {
		log.Fatal("Could not detect supported ESCPOS printer. Check USB connection and try again")
		return
	}

	api := &codeable.ProjectClient{}

	loginRequest := codeable.LoginRequest{
		Email:    os.Getenv("CDBL_EMAIL"),
		Password: os.Getenv("CDBL_PASSWORD"),
	}

	if !loginRequest.IsValid() {
		log.Fatal("Invalid login details provided. Check CDBL_EMAIL | CDBL_PASSWORD env")
		return
	}

	log.Print("Getting login token")
	err = api.Login(loginRequest)
	if err != nil {
		panic(err)
	}

	interval, err := strconv.Atoi(os.Getenv("SCAN_INTERVAL"))
	if interval < 0 || err != nil {
		log.Print("Invalid scan interval set, using default of 5 minutes")
		interval = 5
	}

	for {
		ch := api.StartPoll(interval)
		project := <-ch

		log.Printf("Received new project %d", project.ID)
		fmt.Println(project)
		printProject(printer, project)
	}
}
