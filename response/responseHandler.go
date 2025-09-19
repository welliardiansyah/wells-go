package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"wells-go/application/dtos"
)

type Response struct {
	Success    bool        `json:"success"`
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Timestamp  string      `json:"timestamp"`
	Duration   string      `json:"duration"`
	Data       interface{} `json:"data,omitempty"`
	Error      interface{} `json:"error,omitempty"`
}

type PagingResponse[T any] struct {
	Success    bool   `json:"success"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Timestamp  string `json:"timestamp"`
	Duration   string `json:"duration"`
	Data       []T    `json:"data"`
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
	Total      int64  `json:"total"`
	TotalPages int64  `json:"total_pages"`
}

func localTimeNow() string {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)
	return now.Format("2006-01-02 15:04:05.000")
}

func formatDuration(ms int64) string {
	return fmt.Sprintf("%d ms", ms)
}

func SuccessResponse(w http.ResponseWriter, message string, data interface{}) {
	startedAt := time.Now()
	w.Header().Set("Content-Type", "application/json")
	statusCode := http.StatusOK
	w.WriteHeader(statusCode)

	resp := Response{
		Success:    true,
		StatusCode: statusCode,
		Message:    message,
		Timestamp:  localTimeNow(),
		Duration:   formatDuration(time.Since(startedAt).Milliseconds()),
		Data:       data,
	}

	json.NewEncoder(w).Encode(resp)
}

func SuccessResponsePaging[T any](w http.ResponseWriter, message string, pagingData dtos.PagingResponseFlat[T]) {
	startedAt := time.Now()
	w.Header().Set("Content-Type", "application/json")
	statusCode := http.StatusOK
	w.WriteHeader(statusCode)

	resp := PagingResponse[T]{
		Success:    true,
		StatusCode: statusCode,
		Message:    message,
		Timestamp:  localTimeNow(),
		Duration:   formatDuration(time.Since(startedAt).Milliseconds()),
		Data:       pagingData.Data,
		Page:       pagingData.Page,
		Limit:      pagingData.Limit,
		Total:      pagingData.Total,
		TotalPages: pagingData.TotalPages,
	}

	json.NewEncoder(w).Encode(resp)
}

func ErrorResponse(w http.ResponseWriter, statusCode int, message string, err interface{}) {
	startedAt := time.Now()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	resp := Response{
		Success:    false,
		StatusCode: statusCode,
		Message:    message,
		Timestamp:  localTimeNow(),
		Duration:   formatDuration(time.Since(startedAt).Milliseconds()),
		Error:      err,
	}

	json.NewEncoder(w).Encode(resp)
}
