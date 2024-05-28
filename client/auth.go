package client

import (
	"encoding/json"
	"fmt"
	"shopware6admin/config"
	"shopware6admin/requests"
)

type AuthRequest struct {
	GrantType    string `json:"grant_type"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type AuthResponse struct {
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	AccessToken string `json:"access_token"`
}

type RefreshTokenRequest struct {
	GrantType    string `json:"grant_type"`
	ClientId     string `json:"client_id"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenResponse struct {
	AuthResponse
	RefreshToken string `json:"refresh_token"`
}

type Client struct {
	Config config.Config
}

func (c Client) auth() (AuthResponse, error) {
	url := c.Config.AdminApiUrl + "/api/oauth/token"

	payload, err := json.Marshal(AuthRequest{GrantType: "client_credentials", ClientId: c.Config.ClientId, ClientSecret: c.Config.ClientSecret})
	if err != nil {
		panic(err)
	}

	body, err := requests.Post(url, payload)
	if err != nil {
		panic(err)
	}

	var resBody AuthResponse
	err = json.Unmarshal([]byte(body), &resBody)
	if err != nil {
		panic(err)
	}
	if resBody.AccessToken == "" {
		return resBody, fmt.Errorf("error: %s", string(body))
	}

	return resBody, nil
}

func (c Client) refreshToken() (RefreshTokenResponse, error) {
	url := c.Config.AdminApiUrl + "/api/oauth/token"

	payload, err := json.Marshal(RefreshTokenRequest{GrantType: "client_credentials", ClientId: c.Config.ClientId, RefreshToken: "token"})
	if err != nil {
		panic(err)
	}
	body, err := requests.Post(url, payload)
	if err != nil {
		panic(err)
	}

	var resBody RefreshTokenResponse
	err = json.Unmarshal([]byte(body), &resBody)
	if err != nil {
		panic(err)
	}
	if resBody.AccessToken == "" {
		return resBody, fmt.Errorf("error: %s", string(body))
	}

	return resBody, nil
}

func GetClient(config config.Config) Client {
	var c Client
	c.Config = config
	return c
}

func (c Client) Authorize() string {
	token, err := c.refreshToken()
	if err != nil {
		token, err := c.auth()
		if err != nil {
			panic(err)
		}
		return token.AccessToken
	}
	return token.AccessToken
}
