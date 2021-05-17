package utils

import "net/http"

type Error struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

type Response struct {
	Error Error       `json:"error"`
	Data  interface{} `json:"data"`
}

func ResponseSuccess(w http.ResponseWriter, message string, i interface{}) Response {
	w.Header().Set("Content-Type", "application/json")
	return Response{Error{false, message}, i}
}

func ResponseError(message string) Error {
	return Error{true, message}
}

type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

// Routes routes
type Routes []Route

func AddRoute(path string, method string, handler http.HandlerFunc) Route {
	return Route{path, method, handler}
}
