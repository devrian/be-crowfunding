package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Response is ...
type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

// Meta is ...
type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

// APIResponse is ...
func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	result := Response{
		Meta: meta,
		Data: data,
	}

	return result
}

// FormatValidationError is ...
func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}

// APIResponseErrorEntityProcess is ...
func APIResponseErrorEntityProcess(err string, messageError string) Response {
	errorMessage := gin.H{"errors": err}
	response := APIResponse(messageError, http.StatusUnprocessableEntity, "error", errorMessage)

	return response
}

// APIResponseErrorEntityValidation is ...
func APIResponseErrorEntityValidation(err error, messageError string) Response {
	errors := FormatValidationError(err)
	errorMessage := gin.H{"errors": errors}
	response := APIResponse(messageError, http.StatusUnprocessableEntity, "error", errorMessage)

	return response
}
