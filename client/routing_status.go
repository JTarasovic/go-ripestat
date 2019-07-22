package client

const (
	routingStatusPath = "routing-status/data.json"
)

// RoutingStatus GETs a summary of the current BGP routing state of a given IP prefix or ASN, as observed by the RIS route collectors.
// Historical lookups are supported - a query has to be aligned to the times (00:00, 08:00 and 16:00 UTC) when RIS data has been collected.
func (c *Client) RoutingStatus(resource, timestamp, minPeersSeeing string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = resource
	m["timestamp"] = timestamp
	m["min_peers_seeing"] = minPeersSeeing
	return c.execute(routingStatusPath, m)
}
