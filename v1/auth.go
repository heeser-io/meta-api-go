package v1

import (
	"context"
	"errors"
)

type AuthClient struct {
	client *Client
}

type GenerateTokenParams struct {
	UserID    string                 `json:"userId"`
	ProjectID string                 `json:"projectId"`
	Meta      map[string]interface{} `json:"meta"`
	Auth      map[string]interface{} `json:"auth"`
}

// GenerateToken will generate a token
func (ac *AuthClient) GenerateToken(params *GenerateTokenParams) (string, error) {
	req, err := ac.client.NewRequest(context.Background(), "POST", "socket/create-token", structToReader(params))
	if err != nil {
		return "", err
	}

	var token interface{}
	_, err = ac.client.Do(req, &token)
	if err != nil {
		return "", err
	}
	r, ok := token.(map[string]interface{})
	if !ok {
		return "", errors.New("cannot parse result")
	}
	return r["accessToken"].(string), nil
}
