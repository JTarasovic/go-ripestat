package client

const (
	dnsChainPath = "dns-chain/data.json"
)

// DNSChain GETs the recursive chain of DNS forward (A/AAAA/CNAME) and reverse (PTR) records starting form either a hostname or an IP address.
func (c *Client) DNSChain(resource string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = resource
	return c.execute(dnsChainPath, m)
}
