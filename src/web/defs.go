package main

// ApiBody for request to api
type ApiBody struct {
	// Url for url
	Url     string `json:"url"`
	Method  string `json:"method"`
	ReqBody string `json:"req_body"`
}

type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

var (
	ErrorRequestNotRecognized   = Err{Error: "Api not recognized, bad request", ErrorCode: "001"}
	ErrorRequestBodyParseFailed = Err{Error: "Request body is not correct", ErrorCode: "002"}
	ErrorInternalFaults         = Err{Error: "Internal server error", ErrorCode: "003"}
)
