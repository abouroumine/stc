package api_server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func PrepareResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
}
