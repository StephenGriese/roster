package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/roster", mainPage)
	mux.HandleFunc("/roster/players-for-team", playersForTeam)
	mux.Handle("/*", http.FileServer(http.Dir("./web/static/")))

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
