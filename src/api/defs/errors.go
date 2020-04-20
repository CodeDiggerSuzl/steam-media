package defs

// Err def
type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

// ErrorResponse def
type ErrorResponse struct {
	// HttpSC code
	HttpSC int
	Error  Err
}

var (
	// ErrorRequestBodyParseFailed body
	ErrorRequestBodyParseFailed = ErrorResponse{HttpSC: 400, Error: Err{Error: "Request body is not ok", ErrorCode: "001"}}
	// ErrorNotAuthUser not auth
	ErrorNotAuthUser = ErrorResponse{HttpSC: 401, Error: Err{Error: "Not authorized user", ErrorCode: "002"}}
)
