package api

import (
	"io"
	"net/http"

	"github.com/cadigun/gohttpclient/util"
)

var EmptyResponseBody ResponseBody

// Response struct to hold HTTP response details
type ResponseBody struct {
	resource *http.Response
}

func ResourceToResponseBody(resource *http.Response) ResponseBody {
	return ResponseBody{resource: resource}
}

func (r *ResponseBody) GetResponse() *http.Response {
	return r.resource
}

func (r *ResponseBody) GetStatusCode() int {
	return r.resource.StatusCode
}

func (r *ResponseBody) GetParsedByte() ([]byte, error) {
	defer r.resource.Body.Close()

	return io.ReadAll(r.resource.Body)
}

func (r *ResponseBody) Unmarshall(v any) error {
	data, err := r.GetParsedByte()
	if err != nil {
		return err
	}
	return util.DecodeByteToJson(data, v)
}

type RequestBody struct {
	Headers map[string]string
	URL     string
	Payload interface{}
}
