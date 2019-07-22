package client

const (
	rirPrefixSizeDistributionPath = "rir-prefix-size-distribution/data.json"
)

// RIRPrefixSizeDistribution GETs  the number of allocations and assignments (below the queried resource) according to registration data provided by Regional Internet Registries.
func (c *Client) RIRPrefixSizeDistribution(resource, queryTime string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = resource
	m["queryTime"] = resource
	return c.execute(rirPrefixSizeDistributionPath, m)
}
