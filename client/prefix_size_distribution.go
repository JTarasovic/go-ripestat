package client

const (
	prefixSizeDistributionPath = "prefix-size-distribution/data.json"
)

// PrefixSizeDistribution GETs the total amount of prefixes announced by a given ASN per subnet size and IP version.
func (c *Client) PrefixSizeDistribution(resource, timestamp, minPeersSeeing string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = resource
	m["timestamp"] = resource
	m["min_peers_seeing"] = resource
	return c.execute(prefixSizeDistributionPath, m)
}
