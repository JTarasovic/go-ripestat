package client

import (
	"github.com/go-resty/resty/v2"
)

const (
	baseURL        = "http://stat.ripe.net/data/"
	sourceAppParam = "sourceapp"
)

// Client is the RIPEstat client
type Client struct {
	c *resty.Client
}

// NewClient creates a new RIPEstat client
func NewClient() *Client {
	c := resty.New().SetHostURL(baseURL)
	// c.OnBeforeRequest(func(c *resty.Client, req *resty.Request) error {
	// 	fmt.Printf("URL: %s\n", req.URL)
	// 	fmt.Printf("Method: %s\n", req.Method)
	// 	fmt.Printf("Params: %s\n", req.QueryParam)
	// 	return nil
	// })
	return &Client{c: c}
}

// WithSourceApp sets the `sourceapp` parameter for every request made by this client.
func (c *Client) WithSourceApp(s string) *Client {
	c.c.SetQueryParam(sourceAppParam, s)
	return c
}

func (c *Client) execute(s string, params map[string]string) (map[string]interface{}, error) {
	r := c.c.R()
	r.SetQueryParams(cleanMap(params))
	r.SetResult(map[string]interface{}{})
	resp, err := r.Get(s)
	if err != nil {
		return nil, err
	}
	return *resp.Result().(*map[string]interface{}), nil
}

func cleanMap(m map[string]string) map[string]string {
	for k, v := range m {
		if v == "" {
			delete(m, k)
		}
	}
	return m
}
