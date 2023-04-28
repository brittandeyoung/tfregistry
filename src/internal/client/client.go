package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// HostURL - Default Cloud Services Platform URL
const HostURL string = "https://judp069no2.execute-api.us-east-1.amazonaws.com"

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	// Token      string
	// Auth       AuthStruct
}

// AuthStruct -
// type AuthStruct struct {
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// }

// AuthResponse -
// type AuthResponse struct {
// 	UserID   int    `json:"user_id`
// 	Username string `json:"username`
// 	Token    string `json:"token"`
// }

// NewClient -
func NewClient(host string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    HostURL,
	}

	if host != "" {
		c.HostURL = host
	}

	// If username or password not provided, return empty client
	// if username == nil || password == nil {
	// 	return &c, nil
	// }

	// c.Auth = AuthStruct{
	// 	Username: *username,
	// 	Password: *password,
	// }

	// ar, err := c.SignIn()
	// if err != nil {
	// 	return nil, err
	// }

	// c.Token = ar.Token

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	// token := c.Token

	// if authToken != nil {
	// 	token = *authToken
	// }

	// req.Header.Set("Authorization", token)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusNoContent && res.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
