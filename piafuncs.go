package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const regionURL string = "https://serverlist.piaservers.net/vpninfo/servers/v6"
const tokenURL string = "https://www.privateinternetaccess.com/api/client/v2/token"
const contentTypeForm string = "application/x-www-form-urlencoded"


func getRegionData() RegionData {
	regionDataJson := makeGETRequest(regionURL)
	regionDataJson = bytes.Split(regionDataJson, []byte("\n"))[0]
	var regionData RegionData
	json.Unmarshal(regionDataJson, &regionData)
	return regionData
}

func getToken(username, password string) PIAToken {
	
	reqBody := []byte(fmt.Sprintf("username=%s&password=%s", username, password))
	
	tokenJson := makePOSTRequest(
		tokenURL,
		contentTypeForm,
		reqBody,
	)
	var piaToken PIAToken
	json.Unmarshal(tokenJson, &piaToken)
	return piaToken
}