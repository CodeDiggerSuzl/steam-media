package defs

// Err def
type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

// ErrorResponse def
type ErrorResponse struct {
	// HTTPSC code
	HTTPSC int
	Error  Err
}

var (
	// ErrorRequestBodyParseFailed body
	ErrorRequestBodyParseFailed = ErrorResponse{HTTPSC: 400, Error: Err{Error: "Request body is not ok", ErrorCode: "001"}}
	// ErrorNotAuthUser not auth
	ErrorNotAuthUser = ErrorResponse{HTTPSC: 401, Error: Err{Error: "Not authorized user", ErrorCode: "002"}}
	// ErrorDBError error during db ops
	ErrorDBError = ErrorResponse{HTTPSC: 500, Error: Err{Error: "DB ops failed", ErrorCode: "003"}}
	// ErrorInternalFaults internal error
	ErrorInternalFaults = ErrorResponse{HTTPSC: 500, Error: Err{Error: "Internal service error", ErrorCode: "004"}}
)
