package helpers

import (
	"encoding/json"
	"net/http"
)

// WriteJsonResponse envia uma resposta JSON padronizada para o cliente
func WriteJsonResponse(w http.ResponseWriter, status int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body)
}