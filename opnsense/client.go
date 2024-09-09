package opnsense

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// OPNSenseClient is the client to interact with the OPNSense API
type OPNSenseClient struct {
	APIUrl    string
	APIKey    string
	APISecret string
}

// NewClient creates a new OPNSenseClient
func NewClient(apiUrl, apiKey, apiSecret string) *OPNSenseClient {
	return &OPNSenseClient{
		APIUrl:    apiUrl,
		APIKey:    apiKey,
		APISecret: apiSecret,
	}
}

// DoRequest sends an HTTP request to the OPNSense API
func (client *OPNSenseClient) DoRequest(method, endpoint string, body interface{}) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", client.APIUrl, endpoint)
	var reqBody []byte
	var err error
	if body != nil {
		reqBody, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(client.APIKey, client.APISecret)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(respBody))
	}

	return respBody, nil
}
