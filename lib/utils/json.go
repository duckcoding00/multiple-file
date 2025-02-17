package utils

import (
	"encoding/json"
	"net/http"
)

type RespErr struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

type RespOk struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var statusMessages = map[int]string{
	http.StatusOK:                  "OK",
	http.StatusCreated:             "Created",
	http.StatusBadRequest:          "Bad Request",
	http.StatusInternalServerError: "Internal Server Error",
	http.StatusMethodNotAllowed:    "Method Not Allowed",
	http.StatusUnauthorized:        "Unauthorized",
	http.StatusForbidden:           "Forbidden",
	http.StatusNotFound:            "Not Found",
}

func writeJSON(w http.ResponseWriter, statusCode int, data interface{}) error {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(data)
}

func WriteErr(w http.ResponseWriter, statusCode int, err error) error {
	message, exists := statusMessages[statusCode]
	if !exists {
		message = "Error"
	}

	return writeJSON(w, statusCode, RespErr{
		Message: message,
		Error:   err.Error(),
	})
}

func WriteOk(w http.ResponseWriter, statusCode int, data interface{}) error {
	message, exists := statusMessages[statusCode]
	if !exists {
		message = "Success"
	}

	return writeJSON(w, statusCode, RespOk{
		Message: message,
		Data:    data,
	})
}
