package client

const (
	abuseContactFinderPath = "abuse-contact-finder/data.json"
)

// AbuseContactFinder GETs abuse contact informations for a Internet number resource. Note that this information is in many cases incorrect or not available.
//
// Additional information:
//     - the resources RIR name (if available)
//     - whether the queries resource is related to a special purpose Internet number resource (if available)
//     - blacklist informations (if available)
//     - additional information about the matchin autnum or inet(6)num object in the RIPE DB (e.g. holder name)
//     - less and more specific IP prefixes/ranges for IP based queries
func (c *Client) AbuseContactFinder(resource string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = resource
	return c.execute(abuseContactFinderPath, m)
}
