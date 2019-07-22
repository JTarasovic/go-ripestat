package client

const (
	rirPath = "rir/data.json"
)

// RIR GETs which RIR(s) allocated/assigned a resource.
// Depending on the level of detail ("lod" parameter) this can include additional information like registration status or country of registration.
// The data is based on RIR stats files, see ftp://ftp.ripe.net/pub/stats/.
func (c *Client) RIR(resource, starttime, endtime, levelOfDetail string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = resource
	m["starttime"] = starttime
	m["endtime"] = endtime
	m["lod"] = levelOfDetail
	return c.execute(rirPath, m)
}
