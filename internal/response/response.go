package response

import (
	"encoding/json"
	"net/http"
	"smart-recommendation/internal/errorhandler"
)

const (
	// response code for failed auth
	FailedAuth = "0001"
)

type responseData struct {
	ResponseCode string
	Message      string
	Data         interface{}
}

func Response(w http.ResponseWriter, status int, data interface{}) {
	dataResponse, err := json.Marshal(data)
	if err != nil {
		w.Header().Add("Content-type", "application/json")
		w.WriteHeader(400)
		_, err := w.Write([]byte("Something wrong when parse data"))
		errorhandler.Errorifpanic(err)
	}

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(dataResponse)
	errorhandler.Errorifpanic(err)
}

func SuccessVersionOne(w http.ResponseWriter) {
	result := responseData{
		ResponseCode: "0000",
		Message:      "Welcome to service Smart Brimo version 1, This service is useful for getting recommendation data from AI, Enjoy And CHEERSS :) ",
	}
	Response(w, http.StatusAccepted, result)
}

func Success(w http.ResponseWriter, data interface{}) {
	result := responseData{
		ResponseCode: "0000",
		Message:      "success",
		Data:         data,
	}
	Response(w, http.StatusAccepted, result)
}

func Failed(w http.ResponseWriter, data errorhandler.ErrorData) {
	result := responseData{
		ResponseCode: data.ResponseCode,
		Message:      data.Message,
	}
	Response(w, http.StatusBadRequest, result)
}
