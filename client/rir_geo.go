package client

const (
	rirGeoPath = "rir-geo/data.json"
)

// RIRGeo GETs geographical information for Internet resources based on RIR Statistics data.
func (c *Client) RIRGeo(resource, queryTime string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = resource
	m["query_time"] = queryTime
	return c.execute(rirGeoPath, m)
}
