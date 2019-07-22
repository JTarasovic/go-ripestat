package client

const (
	bgpUpdateActivityPath = "bgp-update-activity/data.json"
)

// BGPUpdateActivity GETs the number of BGP updates seen over time.
func (c *Client) BGPUpdateActivity(resource, starttime, endtime, maxSamples, minSamplingPeriod, numHours, hideEmptySamples string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = resource
	m["starttime"] = starttime
	m["endtime"] = endtime
	m["max_samples"] = maxSamples
	m["min_sampling_period"] = minSamplingPeriod
	m["num_hours"] = numHours
	m["hide_empty_samples"] = hideEmptySamples
	return c.execute(bgpUpdateActivityPath, m)
}
