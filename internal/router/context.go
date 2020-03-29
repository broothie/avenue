package router

import (
	"context"
	"net/http"
)

func PathVar(r *http.Request, key string) interface{} {
	return getPathVar(r.Context(), key)
}

func PathString(r *http.Request, key string) (string, bool) {
	s, ok := PathVar(r, key).(string)
	return s, ok
}

func PathInt(r *http.Request, key string) (int, bool) {
	i, ok := PathVar(r, key).(int)
	return i, ok
}

type routerKey struct{}

var rKey routerKey

func setPathVar(ctx context.Context, key string, value interface{}) context.Context {
	pathVars, ok := ctx.Value(rKey).(map[string]interface{})
	if !ok {
		pathVars = make(map[string]interface{})
	}

	pathVars[key] = value
	return context.WithValue(ctx, rKey, pathVars)
}

func setPathVarOnRequest(r *http.Request, key string, value interface{}) *http.Request {
	return r.WithContext(setPathVar(r.Context(), key, value))
}

func getPathVar(ctx context.Context, key string) interface{} {
	pathVars, ok := ctx.Value(rKey).(map[string]interface{})
	if !ok {
		return nil
	}

	return pathVars[key]
}
