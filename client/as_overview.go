package client

const (
	asOverviewPath = "as-overview/data.json"
)

// ASOverview GETs general informations about an ASN like its announcement status and the name of its holder according to the WHOIS service.
func (c *Client) ASOverview(as string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = as
	return c.execute(asOverviewPath, m)
}
