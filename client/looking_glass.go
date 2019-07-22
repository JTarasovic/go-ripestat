package client

const (
	lookingGlassPath = "looking-glass/data.json"
)

// LookingGlass GETs information commonly coming from a Looking Glass. The data is based on a data feed from the RIPE NCC's network of BGP route collectors (RIS, see https://www.ripe.net/data-tools/stats/ris for more information).
// The data processing usually happens with a small delay and can be considered near real-time.
// The output is structured by collector node (RRC) accompanied by the location and the BGP peers which provide the routing information.
func (c *Client) LookingGlass(resource string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = resource
	return c.execute(lookingGlassPath, m)
}
