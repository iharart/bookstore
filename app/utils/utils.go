package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func StringToUint(val string) uint {
	resUint32, err := strconv.ParseUint(val, 10, 32)
	if err != nil {
		log.Println(err.Error())
	}
	return uint(resUint32)
}

func StringToInt(val string) int {
	res, err := strconv.Atoi(val)
	if err != nil {
		log.Println(err.Error())
	}
	return res
}

func UintToString(val uint) string {
	return strconv.FormatUint(uint64(val), 10)
}

func ErrorCheck(err error) {
	if err != nil {
		log.Println(err.Error())
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
			_, err := w.Write([]byte(err.Error()))
			if err != nil {
				log.Println(err.Error())
			}
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	log.Println(string(response))
	_, err = w.Write([]byte(response))
	if err != nil {
		log.Println(err.Error())
	}
}

func RespondError(w http.ResponseWriter, code int, message string) {
	RespondJSON(false, w, code, map[string]string{"Error": message})
}

type ErrResult struct {
	Error string
}
