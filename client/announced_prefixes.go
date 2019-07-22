package client

const (
	announcedPrefixesPath = "announced-prefixes/data.json"
)

// AnnouncedPrefixes GETs all announced prefixes for a given ASN. The results can be restricted to a specific time period.
func (c *Client) AnnouncedPrefixes(resource, starttime, endtime, minPeersSeeing string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = resource
	m["starttime"] = starttime
	m["endtime"] = endtime
	m["min_peers_seeing"] = minPeersSeeing
	return c.execute(announcedPrefixesPath, m)
}
