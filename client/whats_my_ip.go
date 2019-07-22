package client

const (
	whatsMyIPPath = "whats-my-ip/data.json"
)

// WhatsMyIP GETs the IP address of the requestor.
func (c *Client) WhatsMyIP() (map[string]interface{}, error) {
	return c.execute(whatsMyIPPath, nil)
}
