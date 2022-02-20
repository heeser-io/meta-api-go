package v1

import (
	"context"
	"fmt"
)

type ProjectClient struct {
	client *Client
}

type Project struct {
	ID          string `json:"id"`
	UserID      string `json:"-"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type CreateProjectParams struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
type ReadProjectParams struct {
	ProjectID string `json:"-"`
}
type UpdateProjectParams struct {
	ProjectID   string
	Name        string `json:"name"`
	Description string `json:"description"`
}
type DeleteProjectParams struct {
	ProjectID string
}
type ListProjectParams struct {
}

func (c *ProjectClient) Create(params *CreateProjectParams) (*Project, error) {
	req, err := c.client.NewRequest(context.Background(), "POST", "projects", structToReader(params))
	if err != nil {
		return nil, err
	}

	projectObj := &Project{}
	_, err = c.client.Do(req, projectObj)
	if err != nil {
		return nil, err
	}

	return projectObj, nil
}

func (c *ProjectClient) Read(params *ReadProjectParams) (*Project, error) {
	projectObj := &Project{}

	req, err := c.client.NewRequest(context.Background(), "GET", fmt.Sprintf("projects/%s", params.ProjectID), nil)
	if err != nil {
		return nil, err
	}

	_, err = c.client.Do(req, projectObj)
	if err != nil {
		return nil, err
	}
	return projectObj, nil
}

func (c *ProjectClient) Update(params *UpdateProjectParams) (*Project, error) {
	projectObj := &Project{}

	req, err := c.client.NewRequest(context.Background(), "PUT", fmt.Sprintf("projects/%s", params.ProjectID), structToReader(params))
	if err != nil {
		return nil, err
	}

	_, err = c.client.Do(req, projectObj)
	if err != nil {
		return nil, err
	}
	return projectObj, nil
}

func (c *ProjectClient) Delete(params *DeleteProjectParams) error {
	req, err := c.client.NewRequest(context.Background(), "DELETE", fmt.Sprintf("projects/%s", params.ProjectID), nil)
	if err != nil {
		return err
	}

	_, err = c.client.Do(req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *ProjectClient) List(params *ListProjectParams) ([]*Project, error) {
	projects := []*Project{}

	req, err := c.client.NewRequest(context.Background(), "GET", "projects", structToReader(params))
	if err != nil {
		return nil, err
	}

	_, err = c.client.Do(req, &projects)
	if err != nil {
		return nil, err
	}

	return projects, nil
}
