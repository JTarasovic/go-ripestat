package client

const (
	blacklistPath = "blacklist/data.json"
)

// Blacklist GETs blacklist related data for a queried resource.
func (c *Client) Blacklist(resource, starttime, endtime string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = resource
	m["starttime"] = starttime
	m["endtime"] = endtime
	return c.execute(blacklistPath, m)
}
