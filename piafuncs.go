package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"sort"
)


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
func getToken(username, password string) (PIAToken, error) {
	// Build the application/x-www-form-urlencoded request body, URL escaping any special characters.
	reqBody := []byte(fmt.Sprintf("username=%s&password=%s", url.QueryEscape(username), url.QueryEscape(password)))

	tokenJson := makePOSTRequest(
		TOKEN_URL,
		CT_FORM,
		reqBody,
	)
	var piaToken PIAToken
	json.Unmarshal(tokenJson, &piaToken)
	if piaToken.Token != "" {
		return piaToken, nil
	}
	return piaToken, errors.New("no token received")
}

// adds generated pubkey to the server, responds with server pubkey, status code, DNS servers etc 
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
