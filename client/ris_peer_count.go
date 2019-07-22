package client

const (
	risPeerCountPath = "ris-peer-count/data.json"
)

// RISPeerCount GETs information on the number of peers as seen by the RIS system.
// The data call supports history and each data point is aligned to the RIS RIB dump times (every 8 hours starting from midnight each day).
// Additionally the data shows the number of full-table peers with paramters to change the threshold (per address family).
func (c *Client) RISPeerCount(starttime, endtime, v4FullPrefixThreshold, v6FullPrefixThreshold string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["starttime"] = starttime
	m["endtime"] = endtime
	m["v4_full_prefix_threshold"] = v4FullPrefixThreshold
	m["v6_full_prefix_threshold"] = v6FullPrefixThreshold
	return c.execute(risPeerCountPath, m)
}
