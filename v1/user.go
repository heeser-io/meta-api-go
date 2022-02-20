package v1

import "context"

type UserClient struct {
	client *Client
}

type User struct {
	ID        string `json:"id" dynamo:"id"`
	Username  string `json:"username" dynamo:"username"`
	Email     string `json:"email" dynamo:"email"`
	Plan      *Plan  `json:"plan" dynamo:"plan"`
	CreatedAt string `json:"createdAt" dynamo:"createdAt"`
	UpdatedAt string `json:"updatedAt" dynamo:"updatedAt"`
}

type Plan struct {
	Type        string `json:"type" dynamo:"type"`
	PriceID     string `json:"priceId" dynamo:"priceId"`
	PurchasedAt string `json:"purchasedAt" dynamo:"purchasedAt"`
	ValidUntil  string `json:"validUntil" dynamo:"validUntil"`
}

type CreateUserParams struct {
	Email string
}
type ReadUserParams struct {
	UserID string
}
type UpdateUserParams struct {
}
type DeleteUserParams struct {
}
type ListUserParams struct {
}

func (c *UserClient) Create(params *CreateUserParams) (*User, error) {
	userObj := &User{}
	req, err := c.client.NewRequest(context.Background(), "POST", "users", structToReader(params))
	if err != nil {
		return nil, err
	}

	_, err = c.client.Do(req, userObj)
	if err != nil {
		return nil, err
	}
	return userObj, nil
}

func (c *UserClient) ReadAuth() (*User, error) {
	userObj := &User{}
	req, err := c.client.NewRequest(context.Background(), "GET", "users", nil)
	if err != nil {
		return nil, err
	}

	_, err = c.client.Do(req, userObj)
	if err != nil {
		return nil, err
	}
	return userObj, nil
}

func (c *UserClient) Read(params *ReadUserParams) (*User, error) {
	return nil, nil
}
