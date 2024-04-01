package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/StephenGriese/roster/handlers"
)

func main() {

	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if addr == ":" {
		addr = ":8080"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/roster", handlers.RosterHandler)

	log.Printf("listening on %s\n", addr)

	log.Fatal(http.ListenAndServe(addr, mux))
}
