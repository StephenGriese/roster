package main

import (
	"github.com/StephenGriese/roster/roster"
	"log"
	"net/http"
)

func mainPage(w http.ResponseWriter, r *http.Request) {
	log.Println("mainPage")
	players := []roster.Player{
		{SweaterNumber: 1, LastName: "Smith", FirstName: "John", Position: roster.Goalie},
		{SweaterNumber: 2, LastName: "Jones", FirstName: "Tom", Position: roster.Defense},
	}
	Page(TeamSelect(), Table(players)).Render(w)
}

func playersForTeam(w http.ResponseWriter, r *http.Request) {
	log.Println("playersForTeam")
	players := []roster.Player{
		{SweaterNumber: 3, LastName: "Johnson", FirstName: "Bill", Position: roster.Forward},
		{SweaterNumber: 4, LastName: "Jackson", FirstName: "Jim", Position: roster.Defense},
	}
	Table(players).Render(w)
}
