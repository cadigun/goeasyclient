package api

var EmptyResponseBody ResponseBody

// Response struct to hold HTTP response details
type ResponseBody struct {
	StatusCode int
	Content    []byte
}

type RequestBody struct {
	Headers map[string]string
	URL     string
	Payload interface{}
}
