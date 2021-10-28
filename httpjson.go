// Package httpjson allows you to respond with JSON to the client.
package httpjson

import (
	"encoding/json"
	"net/http"
)

func writeJSON(w http.ResponseWriter, v interface{}, status int) {
	result, e := json.Marshal(v)
	if e != nil {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(status)
	w.Write(result)
}

type responseFn func(http.ResponseWriter, interface{})

// OK writes a JSON object, with a 200 OK response.
func OK(w http.ResponseWriter, v interface{}) {
	writeJSON(w, v, 200)
}

// Accepted writes a JSON object, with a 202 Accepted response.
func Accepted(w http.ResponseWriter, v interface{}) {
	writeJSON(w, v, 202)
}

// BadRequest writes a JSON object, with a 400 Bad Request response.
func BadRequest(w http.ResponseWriter, v interface{}) {
	writeJSON(w, v, 400)
}

// Unauthorized writes a JSON object, with a 401 Unauthorized response.
func Unauthorized(w http.ResponseWriter, v interface{}) {
	writeJSON(w, v, 401)
}

// Forbidden writes a JSON object, with a 403 Forbidden response.
func Forbidden(w http.ResponseWriter, v interface{}) {
	writeJSON(w, v, 403)
}

// NotFound writes a JSON object, with a 404 Not Found response.
func NotFound(w http.ResponseWriter, v interface{}) {
	writeJSON(w, v, 404)
}

// MethodNotAllowed writes a JSON object, with a 405 Method Not Allowed
// response.
func MethodNotAllowed(w http.ResponseWriter, v interface{}) {
	writeJSON(w, v, 405)
}

// InternalServerError writes a JSOn object, with a 500 Internal Server Error
// response.
func InternalServerError(w http.ResponseWriter, v interface{}) {
	writeJSON(w, v, 500)
}
