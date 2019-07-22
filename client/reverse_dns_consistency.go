package client

const (
	reverseDNSConsistencyPath = "reverse-dns-consistency/data.json"
)

// ReverseDNSConsistency GETs the reverse DNS delegations and its consistency with routed and registered IP space.
// The input can be a single prefix or an ASN, in which case all routed and registered prefixes for this ASN are used as an input.
func (c *Client) ReverseDNSConsistency(resource, ipv4, ipv6 string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = resource
	m["ipv4"] = ipv4
	m["ipv6"] = ipv6
	return c.execute(reverseDNSConsistencyPath, m)
}
