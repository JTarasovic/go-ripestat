package client

const (
	asPathLengthPath = "as-path-length/data.json"
)

// ASPathLength GETs AS-path metrics (e.g. shortest or longest AS-path to other ASNs we are peering with) for the queried ASN.
func (c *Client) ASPathLength(as, sortBy string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = as
	m["sort_by"] = sortBy
	return c.execute(asPathLengthPath, m)
}
