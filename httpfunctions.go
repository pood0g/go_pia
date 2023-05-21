package main

import (
	// "fmt"
	"bytes"
	"io"
	"log"
	"io/ioutil"
	"net/http"
	"crypto/x509"
	"crypto/tls"
	// "os/exec"
)

func handleFatal(err error) {
	if err != nil {
		log.Fatalf("%s", err)
	}
}

func makeGETRequest(url string) []byte {

	resp, err := http.Get(url)
	handleFatal(err)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	handleFatal(err)

	return body
}

func makePOSTRequest(url, contentType string, body []byte) []byte {
	reqBody := bytes.NewReader(body)
	resp, err := http.Post(url, contentType, reqBody)
	handleFatal(err)
	defer resp.Body.Close()
	respBody , err := io.ReadAll(resp.Body)
	handleFatal(err)

	return respBody
}

// func runShellCommand(command string, args []string) {
// 	defer waitGroup.Done()
// 	cmd, err:= exec.Command(command, args...).Output()
// 	handleFatal(err)

// 	fmt.Printf("%s", cmd)
// }

func getPIAConfig(url string, body []byte, client http.Client) []byte {
	reqBody := bytes.NewReader([]byte(body))
	req, err := http.NewRequest(http.MethodGet, url, reqBody)
	handleFatal(err)
	resp, err := client.Do(req)
	handleFatal(err)
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)

	return respBody
}

func getTLSClient(certFile string) http.Client {

	rootCAs := x509.NewCertPool()

	cert, err := ioutil.ReadFile(certFile)
	handleFatal(err)
	if ok := rootCAs.AppendCertsFromPEM(cert); ! ok {
		log.Fatalln("Certificate not added.")
	}

	config := &tls.Config{RootCAs: rootCAs}
	transport := &http.Transport{TLSClientConfig: config}
	client := &http.Client{Transport: transport}

	return *client
}