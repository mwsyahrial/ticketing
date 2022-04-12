package helper

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}

func FormatValidationError(err error) []string {
	var errors []string

	err, ok := err.(validator.ValidationErrors)
	if(ok){
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
		}
	}

	return errors
}

func StringToDate(value string) time.Time {
	layoutFormat := "2006-01-02 15:04 MST"
	var date time.Time
	date, _ = time.Parse(layoutFormat, value)

	return date
}

func DateToString(date time.Time) string {
	layoutFormat := "2006-01-02 15:04 MST"
	value := date.Format(layoutFormat)

	return value
}