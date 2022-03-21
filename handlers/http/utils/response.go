package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func RenderJson(ctx context.Context, w http.ResponseWriter, r *http.Request, response interface{}) {
	var buf bytes.Buffer

	enc := json.NewEncoder(&buf)
	if err := enc.Encode(&response); err != nil {
		log.Fatal(err)
	}
	_, err := fmt.Fprint(w, buf.String())
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
}

func RenderBadRequest(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
}

func RenderUnauthorized(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusUnauthorized)
}

func RenderInternalServerError(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}
