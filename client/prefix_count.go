package client

const (
	prefixCountPath = "prefix-count/data.json"
)

// PrefixCount GETs the number of prefixes announced by a given ASN over time.
func (c *Client) PrefixCount(resource, starttime, endtime, minPeersSeeing, resolution string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = resource
	m["starttime"] = starttime
	m["endtime"] = endtime
	m["min_peers_seeing"] = minPeersSeeing
	m["resolution"] = resolution
	return c.execute(prefixCountPath, m)
}
