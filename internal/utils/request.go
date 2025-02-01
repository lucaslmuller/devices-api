package utils

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type requestIDKey string

const (
	REQUEST_ID_KEY requestIDKey = "x-request-id"
)

func ParseRequest[T any](request *http.Request, params []string) (ctx *context.Context, input *T, paramsMap map[string]string, err error) {
	input = new(T)
	paramsMap = map[string]string{}

	for _, p := range params {
		paramsMap[p] = chi.URLParam(request, p)
	}

	requestID := request.Header.Get("x-request-id")

	if requestID == "" {
		requestID = uuid.NewString()
	}

	c := context.WithValue(request.Context(), REQUEST_ID_KEY, requestID)

	if request.Method == "GET" {
		return &c, input, paramsMap, nil
	}

	if ct := request.Header.Get("Content-Type"); strings.Contains(ct, "application/json") {
		body, err := io.ReadAll(request.Body)
		if err != nil {
			return &c, nil, paramsMap, errors.New("invalid request body")
		}

		defer request.Body.Close()
		err = json.Unmarshal(body, input)

		if err != nil {
			return &c, nil, paramsMap, errors.New("invalid request body")
		}
	}

	return &c, input, paramsMap, nil
}
