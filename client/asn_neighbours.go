package client

const (
	aSNNeighboursPath = "asn-neighbours/data.json"
)

// ASNNeighbours GETs nformation on the network neighbours for a given ASN.
// This includes statistical information as well as the list of observed ASN neighbours.
func (c *Client) ASNNeighbours(resource, starttime string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = resource
	m["starttime"] = starttime
	return c.execute(aSNNeighboursPath, m)
}
