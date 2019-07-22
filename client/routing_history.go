package client

const (
	routingHistoryPath = "routing-history/data.json"
)

// RoutingHistory GETs the history of announcements for prefixes, including the origin ASN and the first hop.
// The data comes from the RIS route collectors.
func (c *Client) RoutingHistory(resource, maxRows, includeFirstHop, normaliseVisibility, minPeers, starttime, endtime string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = resource
	m["max_rows"] = maxRows
	m["include_first_hop"] = includeFirstHop
	m["normalise_visibility"] = normaliseVisibility
	m["min_peers"] = minPeers
	m["starttime"] = starttime
	m["endtime"] = endtime
	return c.execute(routingHistoryPath, m)
}
