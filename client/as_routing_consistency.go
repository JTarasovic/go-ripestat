package client

const (
	aSRoutingConsistencyPath = "as-routing-consistency/data.json"
)

// ASRoutingConsistency GETs the difference between the routing and the registration state of an ASN.
// A filter for BGP routes is applied removing non-globally visible prefixes that are not seen by at least 10 RIS full-table peers.
func (c *Client) ASRoutingConsistency(resource string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = resource
	return c.execute(aSRoutingConsistencyPath, m)
}
