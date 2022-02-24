package v1

import (
	"context"
	"fmt"
)

type ApiKeyClient struct {
	client *Client
}

type ApiKey struct {
	ID          string   `json:"id"`
	ProjectID   string   `json:"projectId"`
	ApiKey      string   `json:"apiKey"`
	UserID      string   `json:"-"`
	Scope       []string `json:"scope"`
	Description string   `json:"description"`
	Typename    string   `json:"-"`
	CreatedAt   string   `json:"createdAt"`
	UpdatedAt   string   `json:"updatedAt"`
}

type CreateApiKeyParams struct {
	ProjectID   string   `json:"-" validate:"required"`
	Scope       []string `json:"scope" validate:"required"`
	Description string   `json:"description"`
}
type ReadApiKeyParams struct {
	ProjectID string `json:"-" validate:"required"`
	ApiKeyID  string `json:"-" validate:"required"`
}
type UpdateApiKeyParams struct {
	ProjectID   string `json:"-" validate:"required"`
	ApiKeyID    string `json:"-" validate:"required"`
	Description string `json:"description" update:"true"`
}
type DeleteApiKeyParams struct {
	ProjectID string `json:"-" validate:"required"`
	ApiKeyID  string `json:"-" validate:"required"`
}
type ListApiKeyParams struct {
	ProjectID string `json:"-" validate:"required"`
}

func (c *ApiKeyClient) Create(params *CreateApiKeyParams) (*ApiKey, error) {
	req, err := c.client.NewRequest(context.Background(), "POST", fmt.Sprintf("keys/%s", params.ProjectID), structToReader(params))
	if err != nil {
		return nil, err
	}

	keyObj := &ApiKey{}
	_, err = c.client.Do(req, keyObj)
	if err != nil {
		return nil, err
	}

	return keyObj, nil
}

func (c *ApiKeyClient) Read(params *ReadApiKeyParams) (*ApiKey, error) {
	keyObj := &ApiKey{}

	req, err := c.client.NewRequest(context.Background(), "GET", fmt.Sprintf("keys/%s/%s", params.ProjectID, params.ApiKeyID), nil)
	if err != nil {
		return nil, err
	}

	_, err = c.client.Do(req, keyObj)
	if err != nil {
		return nil, err
	}
	return keyObj, nil
}

func (c *ApiKeyClient) Update(params *UpdateApiKeyParams) (*ApiKey, error) {
	keyObj := &ApiKey{}

	req, err := c.client.NewRequest(context.Background(), "PUT", fmt.Sprintf("keys/%s/%s", params.ProjectID, params.ApiKeyID), structToReader(params))
	if err != nil {
		return nil, err
	}

	_, err = c.client.Do(req, keyObj)
	if err != nil {
		return nil, err
	}
	return keyObj, nil
}

func (c *ApiKeyClient) Delete(params *DeleteApiKeyParams) error {
	req, err := c.client.NewRequest(context.Background(), "DELETE", fmt.Sprintf("keys/%s/%s", params.ProjectID, params.ApiKeyID), nil)
	if err != nil {
		return err
	}

	_, err = c.client.Do(req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *ApiKeyClient) List(params *ListApiKeyParams) ([]*ApiKey, error) {
	projects := []*ApiKey{}

	req, err := c.client.NewRequest(context.Background(), "GET", fmt.Sprintf("keys/%s", params.ProjectID), structToReader(params))
	if err != nil {
		return nil, err
	}

	_, err = c.client.Do(req, &projects)
	if err != nil {
		return nil, err
	}

	return projects, nil
}
