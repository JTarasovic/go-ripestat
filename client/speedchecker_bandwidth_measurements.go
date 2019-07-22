package client

const (
	speedcheckerBandwidthMeasurementsPath = "speedchecker-bandwidth-measurements/data.json"
)

// SpeedcheckerBandwidthMeasurements GETs bandwidth measurement results collected on the Speedchecker platform.
// The bandwith is measured with HTML5 clients (e.g. https://www.broadbandspeedchecker.co.uk) as well as with Speedchecker's mobile applications.
// The unit of measurement is kilobits per second (Kbps).
// Since 2018-05-28 data is synchronized live between Speedchecker and RIPEstat, which means that every measurement done on the Speedchecker platform will show up in this data call after a few seconds.
func (c *Client) SpeedcheckerBandwidthMeasurements(resource, starttime, endtime string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = resource
	m["starttime"] = starttime
	m["endtime"] = endtime
	return c.execute(speedcheckerBandwidthMeasurementsPath, m)
}
