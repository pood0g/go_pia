package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"sort"
)

// String constants used within the application.
const REGION_URL string = "https://serverlist.piaservers.net/vpninfo/servers/v6"
const TOKEN_URL string = "https://www.privateinternetaccess.com/api/client/v2/token"
const CT_FORM string = "application/x-www-form-urlencoded"

// CA file for PIA server trust
const PIA_CERT string = "./ca.rsa.4096.crt"

// function to get and parse region JSON data
func getPIAServerData() RegionData {
	regionDataJson := makeGETRequest(REGION_URL)
	// Remove the junk at the end of the response body.
	regionDataJson = bytes.Split(regionDataJson, []byte("\n"))[0]
	var regionData RegionData
	json.Unmarshal(regionDataJson, &regionData)
	// Sort the regions by name
	sort.Slice(regionData.Regions[:], func(i, j int) bool {
		return regionData.Regions[i].Name < regionData.Regions[j].Name
	})
	return regionData
}

// function to get the auth token from PIA using username and password POST parameters
func getToken(username, password string) PIAToken {
	// Build the application/x-www-form-urlencoded request body, URL escaping any special characters.
	reqBody := []byte(fmt.Sprintf("username=%s&password=%s", url.QueryEscape(username), url.QueryEscape(password)))

	tokenJson := makePOSTRequest(
		TOKEN_URL,
		CT_FORM,
		reqBody,
	)
	var piaToken PIAToken
	json.Unmarshal(tokenJson, &piaToken)
	return piaToken
}

func getPIAConfig(serverip, serverport, token, pubkey string) PIAConfig {

	var piaConfig PIAConfig

	client := getTLSClient()
	url := fmt.Sprintf("https://%s:%s/addKey?pt=%s&pubkey=%s",
		serverip,
		serverport,
		url.QueryEscape(token),
		url.QueryEscape(pubkey),
	)
	resp := makeGETRequestWithCA(url, client)
	json.Unmarshal(resp, &piaConfig)

	return piaConfig
}
