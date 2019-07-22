package client

const (
	geolocationPath = "geoloc/data.json"
)

// Geolocation geolocation information for the given IP space, or for announced IP prefixes in the case of ASNs.
// IPv4 information is based on GeoLite data created by MaxMind, which is Copyright 2008 MaxMind, Inc. All Rights Reserved.
// Please consult MaxMind's license before using this data for non-internal usage.
// For details on the accuracy of this data, please visit MaxMind's product website.
func (c *Client) Geolocation(resource, timestamp string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = resource
	m["timestamp"] = timestamp
	return c.execute(geolocationPath, m)
}
