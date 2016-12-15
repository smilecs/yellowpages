package messages

import (
	"encoding/json"
	"net/http"
)

// Errors

// Errors struct carries a slice of error struct which in turn are error messages that cnfrm with the json+vdn spec
type Errors struct {
	Errors []*Error `json:"errors"`
}

// Error struct carries deatiled error messages spoken of above
type Error struct {
	ID     string `json:"id"`
	Status int    `json:"status"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

// WriteError is a convenience function to write an error struct back to the requester
func WriteError(w http.ResponseWriter, err *Error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.Status)
	json.NewEncoder(w).Encode(Errors{[]*Error{err}})
}

var (
	ErrBadRequest           = &Error{"bad_request", 400, "Bad request", "Request body is not well-formed. It must be JSON."}
	ErrNotAcceptable        = &Error{"not_acceptable", 406, "Not Acceptable", "Accept header must be set to 'application/vnd.api+json'."}
	ErrUnsupportedMediaType = &Error{"unsupported_media_type", 415, "Unsupported Media Type", "Content-Type header must be set to: 'application/vnd.api+json'."}
	ErrInternalServer       = &Error{"internal_server_error", 500, "Internal Server Error", "Something went wrong."}
	ErrNoAuth               = &Error{"unauthorised", 401, "Unauthorised Access", "Not authenticated. Please login."}
	ErrBadToken             = &Error{"unauthorised", 401, "Unauthorised Access", "Not authenticated. Invalid Token."}
)
