package easyhttp

import (
	"encoding/base64"
	"testing"
)

func Test_GenerateBasicAuth(t *testing.T) {
	testCases := []struct {
		name     string
		username string
		password string
		expected string
	}{
		{
			name:     "Basic Case",
			username: "user",
			password: "pass",
			expected: "Basic " + base64.StdEncoding.EncodeToString([]byte("user:pass")),
		},
		{
			name:     "Empty Username",
			username: "",
			password: "pass",
			expected: "Basic " + base64.StdEncoding.EncodeToString([]byte(":pass")),
		},
		{
			name:     "Empty Password",
			username: "user",
			password: "",
			expected: "Basic " + base64.StdEncoding.EncodeToString([]byte("user:")),
		},
		{
			name:     "Empty Username and Password",
			username: "",
			password: "",
			expected: "Basic " + base64.StdEncoding.EncodeToString([]byte(":")),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := GenerateBasicAuth(tc.username, tc.password)
			if result != tc.expected {
				t.Errorf("GenerateBasicAuth(%q, %q) = %q; want %q", tc.username, tc.password, result, tc.expected)
			}
		})
	}
}

/*
type mockClient struct{}

func (m *mockClient) Post(reqbody api.RequestBody) (api.ResponseBody, error) {
	if strings.Contains(reqbody.URL, "invalid_request") {
		return api.EmptyResponseBody, fmt.Errorf("invalid request")
	}
	if strings.Contains(reqbody.URL, "non_200") {
		return api.ResponseBody{StatusCode: 300, Content: []byte(`{"message": "Unexpected error"}`)}, nil
	}
	return api.ResponseBody{StatusCode: 201, Content: []byte(`{"accesstoken": "Success Token"}`)}, nil
}

func Test_GenerateOAuthBearerToken(t *testing.T) {

	testCases := []struct {
		name         string
		url          string
		clientId     string
		clientSecret string
		expected     error
	}{
		{
			name:         "Success",
			url:          "www.example.com",
			clientId:     "user",
			clientSecret: "pass",
		},
		{
			name:         "Non 200",
			url:          "www.example.com/non_200",
			clientId:     "user",
			clientSecret: "pass",
			expected:     fmt.Errorf("Error: response status code 300 with message {\"message\": \"Unexpected error\"}"),
		},
		{
			name:         "Throws Error",
			url:          "www.example.com/invalid_request",
			clientId:     "user",
			clientSecret: "pass",
			expected:     fmt.Errorf("invalid request"),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := generateOAuthBearerToken(tc.url, tc.clientId, tc.clientSecret, &mockClient{})
			assert.Equal(t, err, tc.expected)
		})
	}
}
*/
