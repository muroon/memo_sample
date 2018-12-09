package api

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// ContextKey key for transaction context
type ContextKey string

const (
	resKey = "http.response.key"
)

var responseMap map[string]http.ResponseWriter

func init() {
	responseMap = map[string]http.ResponseWriter{}
}

// setResponseWriter
func setResponseWriter(ctx context.Context, w http.ResponseWriter) context.Context {
	key := generateNewKey()

	ctx = setResKey(ctx, key)

	responseMap[key] = w

	return ctx
}

// getResponseWriter
func getResponseWriter(ctx context.Context) http.ResponseWriter {
	key := getResKey(ctx)
	var res http.ResponseWriter
	var ok bool
	if res, ok = responseMap[key]; !ok {
		panic(fmt.Errorf("http.ResponseWriter is none. key:%s", key))
	}
	return res
}

// setResKey
func setResKey(ctx context.Context, value string) context.Context {
	return context.WithValue(ctx, ContextKey(resKey), value)
}

// deleteResponseWriter
func deleteResponseWriter(ctx context.Context) {
	key := getResKey(ctx)
	if _, ok := responseMap[key]; ok {
		delete(responseMap, key)
	}
}

// getResKey
func getResKey(ctx context.Context) string {
	return getKey(ctx, ContextKey(resKey))
}

// getKey get key
func getKey(ctx context.Context, ctxKey ContextKey) string {
	key, _ := ctx.Value(ctxKey).(string)
	return key
}

// generateNewKey generate key
func generateNewKey() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%d", rand.Int())
}
