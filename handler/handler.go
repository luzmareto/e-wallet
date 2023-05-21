package handler

import (
	"net/http"
)

type Handler struct {
	UserHandler UserHandler
}

func New(userHandler UserHandler) *Handler {
	return &Handler{UserHandler: userHandler}
}

type webResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func responseBadRequest(message string) (int, webResponse) {
	return http.StatusBadRequest, webResponse{
		Code:    http.StatusBadRequest,
		Status:  "Bad Request",
		Message: message,
	}
}

func responseInternalServerError(message string) (int, webResponse) {
	return http.StatusInternalServerError, webResponse{
		Code:    http.StatusInternalServerError,
		Status:  "Internal Server Error",
		Message: message,
	}
}

func responseNotFound(message string) (int, webResponse) {
	return http.StatusNotFound, webResponse{
		Code:    http.StatusNotFound,
		Status:  "Not Found",
		Message: message,
	}
}

func responseOK(message string, data interface{}) (int, webResponse) {
	if data == nil {
		return http.StatusOK, webResponse{
			Code:    http.StatusOK,
			Status:  "OK",
			Message: message,
		}
	}
	return http.StatusOK, webResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: message,
		Data:    data,
	}
}

func responseCreated(message string, data interface{}) (int, webResponse) {
	if data == nil {
		return http.StatusCreated, webResponse{
			Code:    http.StatusCreated,
			Status:  "Created",
			Message: message,
		}
	}
	return http.StatusCreated, webResponse{
		Code:    http.StatusCreated,
		Status:  "Created",
		Message: message,
		Data:    data,
	}
}
