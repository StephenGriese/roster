package faas

import (
	"net/http"

	"github.com/StephenGriese/hello-api/handlers/rest"
)

func Roster(w http.ResponseWriter, r *http.Request) {
	rest.RosterHandler(w, r)
}
