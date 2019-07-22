package client

import "fmt"

func ExampleClient_WhoisObjectLastUpdated() {
	c := NewClient().WithSourceApp("go-ripestat")
	object := "AS64996"     // the exact object to query for
	objectType := "aut-num" // aut-num, inetnum, person, etc
	source := "RIPE"        // RIPE or APNIC
	timestamp := ""
	compareWithLive := "" // When True (default), the version at the last changed time will be compared with the current live object and indicate if it's different.
	r, err := c.WhoisObjectLastUpdated(object, objectType, source, timestamp, compareWithLive)
	if err != nil {
		panic(err)
	}
	data := r["data"].(map[string]interface{})
	delete(data, "query_time")
	fmt.Printf("%q", data)
	// Output:
	// map["last_updated":<nil> "object":"AS64996" "same_as_live":"yes"]
}
