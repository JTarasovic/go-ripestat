package client

const (
	bgpStatePath = "bgp-state/data.json"
)

// BGPState GETs the state of BGP routes for a resource at a certain point in time, as observed by all the RIS collectors.
// This is derived by applying a computation of state to the RIB dump (granularity=8h) that occurred exactly before that time, using the BGP updates observed between the RIB time and the query time.
func (c *Client) BGPState(resource, timestamp, rrcs, unixtimestamps string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = resource
	m["timestamp"] = timestamp
	m["rrcs"] = rrcs
	m["unix_timestamps"] = unixtimestamps
	return c.execute(bgpStatePath, m)
}
