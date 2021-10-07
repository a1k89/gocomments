package utils

import (
	"encoding/json"
	"net/http"
	"os"
)

var SecretChecker = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		systemSecret := os.Getenv("SECRET")
		secret := r.Header.Get("secret")
		if secret != systemSecret {
			result := map[string]interface{}{}
			result["error"] = "Secret key nod provided"
			response := Message(result)
			w.WriteHeader(http.StatusForbidden)
			Response(w,response)
			return
		}
		next.ServeHTTP(w, r)
	})
}

type ResponseWrapper struct {
	results map[string]interface{}
	error map[string]interface{}
}

func Message(data interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	result["data"] = data

	return result
}

func Response(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		return
	}
}
