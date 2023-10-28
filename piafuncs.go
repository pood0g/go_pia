package main

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/url"
	"sort"
	"time"
)

// function to get and parse region JSON data
func getPIAServerData() (RegionData, error) {
	var regionData RegionData
	var pfRegions []Region

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

func pickRegion(data *RegionData) Region {

	var choice uint8

	fmt.Printf("Available regions:\n\n")
	for i, p := range data.Regions {
		fmt.Printf("\t %s[%d]\t%s %s\n", GREEN, i, RESET, p.Name)

	}

	fmt.Printf("\nPick a Region: ")
	fmt.Scanln(&choice)
	return data.Regions[choice]
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

func refreshPortForward(paySig PIAPayloadAndSignature, config *PIAConfig) {
	defer waitGroup.Done()
	for {
		pfStatus, err := requestBindPort(
			config.ServerVIP,
			"19999",
			paySig,
		)
		if err != nil {
			restartServices()
		}
		if pfStatus.Status == "OK" {
			logInfo("Port Forwarding: " + pfStatus.Message)
		}
		time.Sleep(time.Minute*14 + time.Second*50)
	}
}

func connectToPIA(config *goPiaConfig, region *Region, serverData *RegionData) (PIAConfig, PIAToken) {

	rand_server := rand.Intn(len(region.Servers.Wg))
	ip := region.Servers.Wg[rand_server].IP

	logInfo("Creating WireGuard Key Pair")
	keyPair := genKeyPair()

	// Begin connect to PIA
	logInfo(fmt.Sprintf("Connecting to %s - %s\n", region.Name, ip))
	auth, err := getToken(config.PiaUser, config.PiaPass)
	if err != nil {
		logFatal(err.Error())
	}
	logInfo("Got auth token.")

	piaConfig, err := getPIAConfig(
		ip,
		fmt.Sprintf("%d", serverData.Groups.Wg[0].Ports[0]),
		auth.Token,
		keyPair.pubKey,
	)
	if err != nil {
		logFatal(err.Error())
	}

	logInfo(fmt.Sprintf("Server status %s", piaConfig.Status))

	if piaConfig.Status == "OK" {
		logInfo("Got server config successfully.")
		configFile := genWgConfigFile(piaConfig, keyPair)
		writeFile("/etc/wireguard/pia.conf", configFile)
		logInfo("Bringing up wg interface")
		err := runShellCommand("wg-quick", []string{"up", "pia"})
		if err != nil {
			logFatal(err.Error())
		}
		logInfo("WireGuard connection established")
	} else {
		logFatal("WireGuard connection failed")
	}

	return piaConfig, auth
}
