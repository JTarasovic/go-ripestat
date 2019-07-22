package client

const (
	reverseDNSPath = "reverse-dns/data.json"
)

// ReverseDNS GETs the details of reverse DNS delegations for IP prefixes in the RIPE region.
func (c *Client) ReverseDNS(resource string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = resource
	return c.execute(reverseDNSPath, m)
}
