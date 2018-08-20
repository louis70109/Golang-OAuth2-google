package Controller

import (
	"net/http"
	"encoding/json"
)

func Callback(w http.ResponseWriter, r *http.Request) {
	//header := r.Header.Get("Token")
	res := map[string]string{"Result": "true"}
	response, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
