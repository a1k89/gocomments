package utils

import (
	"encoding/json"
	"net/http"
)

func Message() map[string]interface{} {
	return map[string]interface{}{}
}

func Response(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		return
	}
}
