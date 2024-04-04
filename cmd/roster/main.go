package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/StephenGriese/roster/handlers"
	"github.com/StephenGriese/roster/views"
)

func main() {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Debug("Debug message")
	logger.Info("Info message")
	logger.Warn("Warning message")
	logger.Error("Error message")

	wd, _ := os.Getwd()
	logger.Info("Starting server", "working dir", wd)
	log.Printf("working dir: %s\n", wd)
	fmt.Println("printf working dir: ", wd)

	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if addr == ":" {
		addr = ":8080"
	}

	players := views.NewView("bootstrap", "web/views/contacts.gohtml")

	mux := http.NewServeMux()

	mux.Handle("/*", http.FileServer(http.Dir("./web/static/")))

	mux.HandleFunc("/roster", handlers.HandleGetRoster(logger, players))

	log.Printf("sjg listening on %s\n", addr)

	log.Fatal(http.ListenAndServe(addr, mux))
}
