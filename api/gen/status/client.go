// Code generated by goa v3.15.1, DO NOT EDIT.
//
// status client
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package status

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "status" service client.
type Client struct {
	StatusEndpoint goa.Endpoint
}

// NewClient initializes a "status" service client given the endpoints.
func NewClient(status goa.Endpoint) *Client {
	return &Client{
		StatusEndpoint: status,
	}
}

// Status calls the "Status" endpoint of the "status" service.
func (c *Client) Status(ctx context.Context) (res *StatusResult, err error) {
	var ires any
	ires, err = c.StatusEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(*StatusResult), nil
}
