package client

const (
	reverseDNSIPPath = "reverse-dns-ip/data.json"
)

// ReverseDNSIP GETs the reverse DNS info against a single IP address.
func (c *Client) ReverseDNSIP(resource string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = resource
	return c.execute(reverseDNSIPPath, m)
}
