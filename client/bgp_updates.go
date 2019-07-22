package client

const (
	bgpUpdatesPath = "bgp-updates/data.json"
)

// BGPUpdates GETs the BGP updates observed for a resource over a certain period of time.
func (c *Client) BGPUpdates(resource, starttime, endtime, rrcs, unixtimestamps string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = resource
	m["starttime"] = starttime
	m["endtime"] = endtime
	m["rrcs"] = rrcs
	m["unix_timestamps"] = unixtimestamps
	return c.execute(bgpUpdatesPath, m)
}
