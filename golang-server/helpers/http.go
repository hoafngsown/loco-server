package helpers

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"rz-server/internal/common/errors/application_error"
)

// Response structure for consistent API responses
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   *ErrorData  `json:"error,omitempty"`
}

// ErrorData defines the structure of error information
type ErrorData struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// ParseJSON parses request body into the provided struct
func ParseJSON(r *http.Request, dst interface{}) error {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if len(body) == 0 {
		return errors.New("empty request body")
	}

	return json.Unmarshal(body, dst)
}

// RespondWithError handles error responses with application error
func RespondWithError(w http.ResponseWriter, statusCode int, err application_error.Error) {
	response := Response{
		Success: false,
		Error: &ErrorData{
			Code:    err.GetKey(),
			Message: err.Error(),
		},
	}

	RespondWithJSON(w, statusCode, response)
}

// RespondWithSimpleError handles error responses without application error
func RespondWithSimpleError(w http.ResponseWriter, statusCode int, code string, message string) {
	response := Response{
		Success: false,
		Error: &ErrorData{
			Code:    code,
			Message: message,
		},
	}

	RespondWithJSON(w, statusCode, response)
}

// RespondWithJSON sends a JSON response with the given status code
func RespondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(statusCode)
	w.Write(data)
}

// RespondWithSuccess sends a success response with the given data
func RespondWithSuccess(w http.ResponseWriter, data interface{}) {
	response := Response{
		Success: true,
		Data:    data,
	}

	RespondWithJSON(w, http.StatusOK, response)
}

// RespondWithCreated sends a 201 Created response with the given data
func RespondWithCreated(w http.ResponseWriter, data interface{}) {
	response := Response{
		Success: true,
		Data:    data,
	}

	RespondWithJSON(w, http.StatusCreated, response)
}

// RespondWithNoContent sends a 204 No Content response
func RespondWithNoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}
