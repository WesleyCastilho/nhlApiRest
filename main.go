package main

import (
	"io"
	"log"
	"nhlApiRest/nhlApi"
	"os"
	"time"
)

func main(){
	now := time.Now()

	rosterFile, err := os.OpenFile("roster.txt",os.O_RDWR | os.O_CREATE ,0666)
	if err != nil {
		log.Fatalf("error open file roster.txt: %v", err)
	}
	defer rosterFile.Close()

	wrt := io.MultiWriter(os.Stdout, rosterFile)
	log.SetOutput(wrt)

	teams, err := nhlApi.GetAllTeams()
	if err != nil {
		log.Fatalf("error while getting all teams %v: ", err)
	}

	for _, team := range teams {
		log.Println("-------------------------")
		log.Printf("Name %s", team.Name)
		log.Printf("Conferência %s", team.Conference.Name)
		log.Printf("Divisão %s", team.Division.Name)
		log.Printf("Franquia %s", team.Franchise.Teamname)

		log.Println("-------------------------")

	}

	log.Printf("took %v", time.Now().Sub(now).String())

}