package client

const (
	mlabClientsPath = "mlab-clients/data.json"
)

// MlabClients GETs a set of all the hosts within a certain resource for which any network tests occurred.
// The data is based on active host measurements collected by the Measurement Lab platform (M-Lab).
// The measurements are commonly ran using the M-Lab Network Detection Tool (NDT), available as a stand-alone network speed test application, and also included in a popular BitTorrent client.
// Note that due to the nature the data is processed data can be delayed for around two days at the beginning of each month!
// The results published, including host details are covered by the M-Lab acceptable use policy.
func (c *Client) MlabClients(resource, starttime, endtime string) (map[string]interface{}, error) {
	m := make(map[string]string)
	m["resource"] = resource
	m["starttime"] = starttime
	m["endtime"] = endtime
	return c.execute(mlabClientsPath, m)
}
