package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"sort"
)

// String constants used within the application.
const regionURL string = "https://serverlist.piaservers.net/vpninfo/servers/v6"
const tokenURL string = "https://www.privateinternetaccess.com/api/client/v2/token"
const contentTypeForm string = "application/x-www-form-urlencoded"

// function to get and parse region JSON data
func getRegionData() RegionData {
	regionDataJson := makeGETRequest(regionURL)
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
		tokenURL,
		contentTypeForm,
		reqBody,
	)
	var piaToken PIAToken
	json.Unmarshal(tokenJson, &piaToken)
	return piaToken
}