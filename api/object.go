package api

var EmptyResponseBody ResponseBody

// Response struct to hold HTTP response details
type ResponseBody struct {
	StatusCode int
	Content    []byte
}

type RequestBody[T any] struct {
	Headers map[string]string
	URL     string
	Payload T
}
