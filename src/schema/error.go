package schema

// ErrorResponse represents an error received from the Hetzner DNS API
// Message contains a human readable string describing the error.
// Code represents the HTTP status code
type ErrorResponse struct {
	Error struct {
		Message string `json:"message,omitempty"`
		Code    int    `json:"code,omitempty"`
	} `json:"error"`
}
