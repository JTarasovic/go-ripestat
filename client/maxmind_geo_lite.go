package client

const (
	maxmindGeoLitePath = "maxmind-geo-lite/data.json"
)

// MaxmindGeoLite GETs geolocation information for the given IP space based on MaxMind's GeoLite2 data source.
func (c *Client) MaxmindGeoLite(resource, queryTime string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = resource
	m["query_time"] = queryTime
	return c.execute(maxmindGeoLitePath, m)
}
