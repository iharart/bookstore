package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func StringToUint(val string) uint {
	resUint32, err := strconv.ParseUint(val, 2, 32)
	if err != nil {
		panic(err)
	}
	return uint(resUint32)
}

func ErrorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func RespondJSON(shouldBeEmpty bool, w http.ResponseWriter, status int, payload interface{}) {
	var response []byte
	var err error
	if shouldBeEmpty {
		response, _ = json.Marshal(payload)
	} else {
		response, err = json.Marshal(payload)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	log.Println(string(response))
	w.Write([]byte(response))
}

func RespondError(w http.ResponseWriter, code int, message string) {
	RespondJSON(false, w, code, map[string]string{"error": message})
}
