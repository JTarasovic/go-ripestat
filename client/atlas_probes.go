package client

const (
	atlasProbesPath = "atlas-probes/data.json"
)

// AtlasProbes GETs information on the RIPE Atlas probes in an network (ASN), a prefix or a country.
// The information is based on data coming from the RIPE Atlas REST API, https://atlas.ripe.net/docs/api/v2/manual/.
func (c *Client) AtlasProbes(resource string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = resource
	return c.execute(atlasProbesPath, m)
}
