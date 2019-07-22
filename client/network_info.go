package client

const (
	networkInfoPath = "network-info/data.json"
)

// NetworkInfo GETs the containing prefix and announcing ASN of a given IP address.
func (c *Client) NetworkInfo(resource string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = resource
	return c.execute(networkInfoPath, m)
}
