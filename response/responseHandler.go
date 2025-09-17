package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Response struct {
	Success    bool        `json:"success"`
	StatusCode int         `json:"status_code"`
	Status     string      `json:"status,omitempty"`
	Message    string      `json:"message"`
	Timestamp  string      `json:"timestamp"`
	Duration   string      `json:"duration"`
	Data       interface{} `json:"data,omitempty"`
	Error      interface{} `json:"error,omitempty"`
}

type ResponseOneBilling struct {
	Success   bool        `json:"success"`
	Status    int         `json:"status"`
	Timestamp string      `json:"timestamp"`
	Duration  string      `json:"duration"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	Error     interface{} `json:"error,omitempty"`
}

type ResponseWithLogData struct {
	Response
	LogDeposit interface{} `json:"log_deposit,omitempty"`
}

func localTimeNow() string {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)
	return now.Format("2006-01-02 15:04:05.000")
}

func formatDuration(ms int64) string {
	return fmt.Sprintf("%d ms", ms)
}

func SuccessResponse(w http.ResponseWriter, message string, data interface{}, startedAt time.Time) {
	w.Header().Set("Content-Type", "application/json")
	statusCode := http.StatusOK
	w.WriteHeader(statusCode)

	elapsed := time.Since(startedAt).Milliseconds()
	json.NewEncoder(w).Encode(Response{
		Success:    true,
		StatusCode: statusCode,
		Message:    message,
		Timestamp:  localTimeNow(),
		Duration:   formatDuration(elapsed),
		Data:       data,
	})
}

func ErrorResponse(w http.ResponseWriter, statusCode int, message string, err interface{}, startedAt time.Time) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	elapsed := time.Since(startedAt).Milliseconds()
	json.NewEncoder(w).Encode(Response{
		Success:    false,
		StatusCode: statusCode,
		Message:    message,
		Timestamp:  localTimeNow(),
		Duration:   formatDuration(elapsed),
		Error:      err,
	})
}
