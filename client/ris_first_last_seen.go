package client

const (
	risFirstLastSeenPath = "ris-first-last-seen/data.json"
)

// RISFirstLastSeen GETs information on when a prefix or ASN was first and last seen in RIS data.
// The data generally goes back to 2000.
// For the recency of the data you can check the parameter "latest_time", which usually is not more than 8 hours behind real-time.
// The "low_visibility" flag, which can be optionally included, shows if the data point was seen by a low or high number of peers.
func (c *Client) RISFirstLastSeen(resource, include string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = resource
	m["include"] = include
	return c.execute(risFirstLastSeenPath, m)
}
