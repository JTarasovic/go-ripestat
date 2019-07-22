package client

const (
	mlabBandwidthPath = "mlab-bandwidth/data.json"
)

// MlabBandwidth a set of all the measured network bandwidths for a certain resource.
// The data is based on active host measurements collected by the Measurement Lab platform (M-Lab).
// The bandwidth of a host is determined as the maximum network throughput value for all the tests/measurements performed by that host during the specified time period.
// The value of the measurement throughput is computed as the number of octets transmitted between the host and the chosen M-Lab server, divided by their transfer time.
// The measurements are commonly ran using the M-Lab Network Detection Tool (NDT), available as a stand-alone network speed test application, and also included in a popular BitTorrent client.
// Note that due to the nature the data is processed data can be delayed for around two days at the beginning of each month!
// The results published, including host details are covered by the M-Lab acceptable use policy.
func (c *Client) MlabBandwidth(resource, starttime, endtime string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = resource
	m["starttime"] = starttime
	m["endtime"] = endtime
	return c.execute(mlabBandwidthPath, m)
}
