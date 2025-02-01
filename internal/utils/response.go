package utils

import (
	"context"
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

type SuccessResponse[T any] struct {
	Success bool `json:"success"`
	Data    T    `json:"data"`
}

func RespondWithJSON(ctx *context.Context, w http.ResponseWriter, code int, data any) {
	output := SuccessResponse[any]{
		Success: true,
		Data:    data,
	}

	reqID := (*ctx).Value(REQUEST_ID_KEY)
	if reqID != nil {
		w.Header().Set(string(REQUEST_ID_KEY), reqID.(string))
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	w.WriteHeader(code)
	bOutput, _ := json.Marshal(output)
	w.Write(bOutput)
}

func RespondWithError(ctx *context.Context, w http.ResponseWriter, code int, message string) {
	output := ErrorResponse{
		Success: false,
		Error:   message,
	}

	reqID := (*ctx).Value(REQUEST_ID_KEY)
	if reqID != nil {
		w.Header().Set(string(REQUEST_ID_KEY), reqID.(string))
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	w.WriteHeader(code)
	bOutput, _ := json.Marshal(output)
	w.Write(bOutput)
}
