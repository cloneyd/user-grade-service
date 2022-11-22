package v1

import (
	"fmt"
	"log"
	"net/http"
)

func errorAsJSON(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, printErr := fmt.Fprintf(w, `{"error": "%s"}`, err)
	if printErr != nil {
		log.Println(printErr)
	}
}
