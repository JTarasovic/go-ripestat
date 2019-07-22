package client

const (
	prefixRoutingConsistencyPath = "prefix-routing-consistency/data.json"
)

// PrefixRoutingConsistency GETs the given routes (prefix originating from an ASN) between Routing Registries and actual routing behaviour as seen by the RIPE NCC route collectors (RIS).
func (c *Client) PrefixRoutingConsistency(resource string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = resource
	return c.execute(prefixRoutingConsistencyPath, m)
}
