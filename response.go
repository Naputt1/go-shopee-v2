package goshopee

type BaseResponse struct {
	// The identifier for an API request for error tracking
	RequestID string `json:"request_id"`
	// Indicate error type if hit error. Empty if no error happened.
	Error string `json:"error"`
	// Indicate error details if hit error. Empty if no error happened.
	Message string `json:"message"`
	// Indicate waring details if hit waring. Empty if no waring happened.
	Warning string `json:"warning"`
}
