package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func makeGETRequest(url string) ([]byte, error) {

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}

func makePOSTRequest(url, contentType string, body []byte) ([]byte, error) {
	reqBody := bytes.NewReader(body)
	resp, err := http.Post(url, contentType, reqBody)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBody, err
}

func makeGETRequestWithCA(url string, client http.Client) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	fmt.Printf("%s", req.Header)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)

	return respBody, err
}

// func makePOSTRequestWithCA(url string, client http.Client, body []byte) ([]byte, error) {
// 	reqBody := bytes.NewReader(body)
// 	req, err := http.NewRequest(http.MethodPost, url, reqBody)
// 	if err != nil {
// 		return nil, err
// 	}
// 	req.Header.Set("Content-Type", CT_FORM)
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()
// 	respBody, err := io.ReadAll(resp.Body)

// 	return respBody, err
// }

// work around for making https connection to IP with non trusted CA
func getTLSClient() http.Client {

	// skip initial TLS but verify the peer certificate with custom verification function, not 100% tested
	config := &tls.Config{
		InsecureSkipVerify:    true,
		VerifyPeerCertificate: verifyCert,
	}
	transport := &http.Transport{TLSClientConfig: config}
	client := &http.Client{Transport: transport}

	return *client
}

func verifyCert(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error {

	rootCAs := x509.NewCertPool()

	caCert, err := os.ReadFile(PIA_CERT)
	logFatal(err)
	if ok := rootCAs.AppendCertsFromPEM(caCert); !ok {
		log.Fatalln("Certificate not added.")
	}
	log.Printf("Certificate %s parsed successfully\n", PIA_CERT)

	hostCert, _ := x509.ParseCertificate(rawCerts[0])
	opts := x509.VerifyOptions{
		Roots: rootCAs,
	}
	if _, err := hostCert.Verify(opts); err != nil {
		log.Println("Unable to verify cert")
		return err
	}
	log.Println("[+] Server certificate validated.")

	return nil
}
