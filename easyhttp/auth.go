package easyhttp

import (
	"encoding/base64"
	"fmt"

	"github.com/cadigun/goeasyclient/api"
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
	if !(resp.GetStatusCode() >= 200 && resp.GetStatusCode() < 300) {
		data, _ := resp.GetParsedByte()
		return "", fmt.Errorf("Error: response status code %v with message %v", resp.GetStatusCode(), string(data))
	}
	var tokenObj TokenResponse
	err = resp.Unmarshall(&tokenObj)
	if err != nil {
		return "", err
	}
	return "Bearer " + tokenObj.AccessToken, nil
}
