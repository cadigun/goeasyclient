package easyhttp

import "github.com/cadigun/gohttpclient/api"

type EasyHttpClient interface {
	Post(requestbody api.RequestBody) (api.ResponseBody, error)
}
