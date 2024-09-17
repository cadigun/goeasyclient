package easyhttp

import (
	"net/http"

	"github.com/cadigun/goeasyclient/api"
	"github.com/cadigun/goeasyclient/util"
)

var easyhttpDefault = New()

func Default() *EasyHttp {
	return easyhttpDefault
}

var defaultHeaders = map[string]string{
	"Content-Type": "application/json",
}

type EasyHttp struct {
	httpClient *http.Client
}

func New() *EasyHttp {
	return &EasyHttp{httpClient: &http.Client{}}
}

func (c *EasyHttp) Post(requestbody api.RequestBody) (api.ResponseBody, error) {
	return c.doRequestWithCustomHeaders("POST", requestbody.URL, requestbody.Payload, requestbody.Headers)
}

func (c *EasyHttp) Put(requestbody api.RequestBody) (api.ResponseBody, error) {
	return c.doRequestWithCustomHeaders("PUT", requestbody.URL, requestbody.Payload, requestbody.Headers)
}

func (c *EasyHttp) Patch(requestbody api.RequestBody) (api.ResponseBody, error) {
	return c.doRequestWithCustomHeaders("PATCH", requestbody.URL, requestbody.Payload, requestbody.Headers)
}

func (c *EasyHttp) Delete(requestbody api.RequestBody) (api.ResponseBody, error) {
	return c.doRequestWithCustomHeaders("DELETE", requestbody.URL, requestbody.Payload, requestbody.Headers)
}

func (c *EasyHttp) Get(requestbody api.RequestBody) (api.ResponseBody, error) {
	return c.doRequestWithCustomHeaders("GET", requestbody.URL, requestbody.Payload, requestbody.Headers)
}

func (c *EasyHttp) Do(method string, requestbody api.RequestBody) (api.ResponseBody, error) {
	return c.doRequestWithCustomHeaders(method, requestbody.URL, requestbody.Payload, requestbody.Headers)
}

func (c *EasyHttp) doRequest(method, url string, data interface{}) (api.ResponseBody, error) {
	return c.doRequestWithCustomHeaders(method, url, data, nil)
}

func (c *EasyHttp) doRequestWithCustomHeaders(method, url string, data interface{}, customHeaders map[string]string) (api.ResponseBody, error) {
	if customHeaders == nil {
		customHeaders = map[string]string{}
	}
	payload, err := util.EncodeObjectToBytesBuffer(data)
	if err != nil {
		return api.EmptyResponseBody, err
	}

	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return api.EmptyResponseBody, err
	}

	for k, v := range defaultHeaders {
		if _, ok := customHeaders[k]; !ok {
			customHeaders[k] = v
		}
	}

	for k, v := range customHeaders {
		req.Header.Add(k, v)
	}

	response, err := c.httpClient.Do(req)
	if err != nil {
		return api.EmptyResponseBody, err
	}

	return api.ResourceToResponseBody(response), nil
}
