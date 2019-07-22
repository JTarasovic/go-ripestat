package client

const (
	maxmindGeoLiteAnnouncedByASPath = "maxmind-geo-lite-announced-by-as/data.json"
)

// MaxmindGeoLiteAnnouncedByAS GETs geolocation information for prefixes that are announced by an autonomous system. Annoucement information is based on RIS (http://ris.ripe.net), geolocation information is based on MaxMind's GeoLite2 data.
// Prefix information (IPv4/IPv6) is based on GeoLite2 data created by MaxMind, which is Copyright 2008 MaxMind, Inc. All Rights Reserved.
// Please consult MaxMind's license before using this data for non-internal usage.
// For details on the accuracy of this data, please visit MaxMind's product website.
func (c *Client) MaxmindGeoLiteAnnouncedByAS(resource string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = resource
	return c.execute(maxmindGeoLiteAnnouncedByASPath, m)
}
