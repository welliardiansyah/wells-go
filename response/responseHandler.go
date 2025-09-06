package response

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
)

type Response struct {
	Success    bool        `json:"success"`
	StatusCode int         `json:"status_code"`
	Status     string      `json:"status,omitempty"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Error      interface{} `json:"error,omitempty"`
}
type ResponseOneBilling struct {
	Success   bool        `json:"success"`
	Status    int         `json:"status"`
	Timestamp string      `json:"timestamp"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	Error     interface{} `json:"error,omitempty"`
}

type ResponseWithLogData struct {
	Response
	LogDeposit interface{} `json:"log_deposit,omitempty"`
}

func SuccessResponse(w http.ResponseWriter, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	statusCode := http.StatusOK
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(Response{
		Success:    true,
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	})
}
func SuccessResponseOneBilling(w http.ResponseWriter, message, timestamp string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	statusCode := http.StatusOK
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(ResponseOneBilling{
		Success:   true,
		Status:    statusCode,
		Timestamp: timestamp,
		Message:   message,
		Data:      data,
		Error:     nil,
	})
	if err != nil {
		log.Err(err).Msg("error on encode response")
	}
}

func ErrorResponse(w http.ResponseWriter, statusCode int, message string, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(Response{
		Success:    false,
		StatusCode: statusCode,
		Message:    message,
		Error:      err,
	})
}
