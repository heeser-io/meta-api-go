package v1

import "context"

type SocketApiClient struct {
	client *Client
}

func (c *SocketApiClient) DispatchEvent(params *DispatchEventParams) error {
	req, err := c.client.NewRequest(context.Background(), "POST", "socket/dispatch", structToReader(params))
	if err != nil {
		return err
	}

	_, err = c.client.Do(req, &[]byte{})
	if err != nil {
		return err
	}
	return nil
}
