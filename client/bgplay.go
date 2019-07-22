package client

const (
	bgplayPath = "bgplay/data.json"
)

// BGPlay GETs a representation of what occurred to the BGP routes of a resource over a period of time.
// It includes data that defines the initial BGP state at the start time of the query, and all the BGP updates observed from then until the end time, as well as a description of all the AS nodes, and RIS BGP peers involved in the result.
func (c *Client) BGPlay(resource, starttime, endtime, rrcs, unixtimestamps string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = resource
	m["starttime"] = starttime
	m["endtime"] = endtime
	m["rrcs"] = rrcs
	m["unix_timestamps"] = unixtimestamps
	return c.execute(bgplayPath, m)
}
