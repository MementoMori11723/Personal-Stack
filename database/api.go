package database

import (
	"encoding/json"
	"net/http"
)

var (
	apiMux = http.NewServeMux()
)

func init() {
	apiMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"message": "Hello, World!"})
	})
}

func API() *http.ServeMux {
  return apiMux
}