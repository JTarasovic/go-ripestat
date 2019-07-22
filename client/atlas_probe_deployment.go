package client

const (
	atlasProbeDeploymentPath = "atlas-probe-deployment/data.json"
)

// AtlasProbeDeployment GETs information on the number of RIPE Atlas probes in a region, a country or network (ASN).
// It supports history, with a general start in 2014.
// The information is based on data from the RIPE Atlas probe archive, ftp://ftp.ripe.net/ripe/atlas/probes/archive/, which is processed once a day.
func (c *Client) AtlasProbeDeployment(resource, starttime, endtime string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = resource
	m["starttime"] = starttime
	m["endtime"] = endtime
	return c.execute(atlasProbeDeploymentPath, m)
}
