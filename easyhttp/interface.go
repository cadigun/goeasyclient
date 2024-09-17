package easyhttp

import "github.com/cadigun/goeasyclient/api"

type EasyHttpClient interface {
	Post(requestbody api.RequestBody) (api.ResponseBody, error)
}
