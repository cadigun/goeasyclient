package easyhttp

import (
	"encoding/base64"
	"fmt"

	"github.com/cadigun/gohttpclient/api"
	"github.com/cadigun/gohttpclient/util"
)

func GenerateBasicAuth(username, password string) string {
	auth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
	return "Basic " + auth
}

type TokenRequest struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

func GenerateOAuthBearerToken(url, clientId, clientSecret string) (string, error) {
	return generateOAuthBearerToken(url, clientId, clientSecret, nil)
}

func generateOAuthBearerToken(url, clientId, clientSecret string, easyClient EasyHttpClient) (string, error) {
	payload := TokenRequest{
		GrantType:    "client_credentials",
		ClientID:     clientId,
		ClientSecret: clientSecret,
	}
	if easyClient == nil {
		easyClient = Default()
	}
	resp, err := easyClient.Post(api.RequestBody{URL: url, Payload: payload})
	if err != nil {
		return "", err
	}
	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		return "", fmt.Errorf("Error: response status code %v with message %v", resp.StatusCode, string(resp.Content))
	}
	var tokenObj TokenResponse
	err = util.DecodeByteToJson(resp.Content, &tokenObj)
	if err != nil {
		return "", err
	}
	return "Bearer " + tokenObj.AccessToken, nil
}
