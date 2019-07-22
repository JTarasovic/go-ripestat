package client

const (
	zonemasterPath = "zonemaster/data.json"
)

// Zonemaster GETs the test results of DNS checks run by Zonemaster.
// The data call has two modes, one to get an overview of available tests ("resource" parameter is hostname)
// and one to get the test details ("resource" parameter is test ID and "method" parameter is "details").
// Please note that this data call is in development and features/availability can change.
func (c *Client) Zonemaster(resource, method string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = resource
	m["method"] = method
	return c.execute(zonemasterPath, m)
}
