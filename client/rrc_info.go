package client

const (
	rrcInfoPath = "rrc-info/data.json"
)

// RRCInfo GETs information on collector nodes (RRCs) of the RIS network (http://ris.ripe.net).
// This includes geographical and topological location and information on collectors' peers.
func (c *Client) RRCInfo() (map[string]interface{}, error) {
	return c.execute(rrcInfoPath, nil)
}
