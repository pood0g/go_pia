package main

import (
	"encoding/json"
	"bytes"
)

func getRegionData() RegionData {
	regionData := makeGETRequest("https://serverlist.piaservers.net/vpninfo/servers/v6")
	regionData = bytes.Split(regionData, []byte("\n"))[0]
	var jsonData RegionData
	json.Unmarshal(regionData, &jsonData)
	return jsonData
}