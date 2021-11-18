package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	fmt.Println(string(response))
	w.Write([]byte(response))
}

func RespondError(w http.ResponseWriter, code int, message string) {
	RespondJSON(w, code, map[string]string{"error": message})
}

func ErrorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}
