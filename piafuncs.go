package main

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"sort"
)

// function to get and parse region JSON data
func getPIAServerData() (RegionData, error) {
	var regionData RegionData
	var pfRegions []Regions

	regionDataJson, err := makeGETRequest(REGION_URL)
	// Remove the junk at the end of the response body.
	regionDataJson = bytes.Split(regionDataJson, []byte("\n"))[0]
	json.Unmarshal(regionDataJson, &regionData)
	// Remove non port forwarding regions
	for _, region := range regionData.Regions {
		if region.PortForward {
			pfRegions = append(pfRegions, region)
		}
	}
	regionData.Regions = pfRegions
	// Sort the regions by name
	sort.Slice(regionData.Regions[:], func(i, j int) bool {
		return regionData.Regions[i].Name < regionData.Regions[j].Name
	})
	return regionData, err
}

// function to get the auth token from PIA using username and password POST parameters
func getToken(username, password string) (PIAToken, error) {
	// Build the application/x-www-form-urlencoded request body, URL escaping any special characters.
	reqBody := []byte(fmt.Sprintf("username=%s&password=%s", url.QueryEscape(username), url.QueryEscape(password)))

	tokenJson, err := makePOSTRequest(
		TOKEN_URL,
		CT_FORM,
		reqBody,
	)
	var piaToken PIAToken
	json.Unmarshal(tokenJson, &piaToken)
	if piaToken.Token != "" {
		return piaToken, err
	}
	return piaToken, err
}

// adds generated pubkey to the server, responds with server pubkey, status code, DNS servers etc
func getPIAConfig(serverIP, serverPort, token, pubKey string) (PIAConfig, error) {

	var piaConfig PIAConfig

	reqURL := fmt.Sprintf("https://%s:%s/addKey?pt=%s&pubkey=%s",
		serverIP,
		serverPort,
		url.QueryEscape(token),
		url.QueryEscape(pubKey),
	)
	resp, err := makeGETRequestWithCA(reqURL)
	json.Unmarshal(resp, &piaConfig)

	return piaConfig, err
}

func getPFSignature(serverIP, serverPort, token string) (PIAPayloadAndSignature, PIAPFPayload, error) {

	var payloadAndSignature PIAPayloadAndSignature
	var payload PIAPFPayload

	reqURL := fmt.Sprintf("https://%s:%s/getSignature?token=%s",
		serverIP,
		serverPort,
		url.QueryEscape(token),
	)
	resp, err := makeGETRequestWithCA(reqURL)
	json.Unmarshal(resp, &payloadAndSignature)
	payload_json, _ := b64.StdEncoding.DecodeString(payloadAndSignature.Payload)
	json.Unmarshal(payload_json, &payload)

	return payloadAndSignature, payload, err
}

func requestBindPort(serverIP, serverPort string, pldSig PIAPayloadAndSignature) (PIAPFStatus, error) {

	var pfStatus PIAPFStatus

	reqURL := fmt.Sprintf("https://%s:%s/bindPort?payload=%s&signature=%s",
		serverIP,
		serverPort,
		url.QueryEscape(pldSig.Payload),
		url.QueryEscape(pldSig.Signature),
	)
	resp, err := makeGETRequestWithCA(reqURL)
	json.Unmarshal(resp, &pfStatus)

	return pfStatus, err
}
