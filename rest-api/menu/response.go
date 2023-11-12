package menu

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"error,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}

func WriteJsonResponse(w http.ResponseWriter, resp APIResponse) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(resp.Status)

	json.NewEncoder(w).Encode(resp)
}
