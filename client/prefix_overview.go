package client

const (
	prefixOverviewPath = "prefix-overview/data.json"
)

// PrefixOverview GETs a summary of the given prefix, including whether and by whom it is announced.
func (c *Client) PrefixOverview(resource, minPeersSeeing, maxRelated, queryTime string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = resource
	m["min_peers_seeing"] = minPeersSeeing
	m["max_related"] = maxRelated
	m["query_time"] = queryTime
	return c.execute(prefixOverviewPath, m)
}
