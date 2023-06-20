package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/curve25519"
)

// converts []byte to [32]byte
func conv32(key *[]byte) *[32]byte {
	var b32 [32]byte
	copy(b32[:], *key)
	return &b32
}

// converts [32]byte to []byte
func conv(key *[32]byte) *[]byte {
	b := key[:]
	return &b
}

// generates a cryptographically secure private key
func genPrivKey() *[32]byte {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	logFatal(err)
	key[0] &= 248
	key[31] &= 127
	key[31] |= 64

	prvKey := conv32(&key)

	return prvKey
}

// generate the public key from the private key
func genPubkey(privKey *[32]byte) (*[]byte, *[]byte) {
	pubKey := [32]byte{}
	curve25519.ScalarBaseMult(&pubKey, privKey)
	return conv(privKey), conv(&pubKey)
}

func genKeyPair() WGKeyPair {

	prvKey, pubKey := genPubkey(genPrivKey())

	prvKeyb64 := base64.StdEncoding.EncodeToString(*prvKey)
	pubKeyb64 := base64.StdEncoding.EncodeToString(*pubKey)

	return WGKeyPair{
		prvKey: prvKeyb64,
		pubKey: pubKeyb64,
	}
}

func genWgConfigFile(conf PIAConfig, keys WGKeyPair) []byte {

	iface := WgConfigInterface{
		Address:    conf.PeerIP,
		PrivateKey: keys.prvKey,
		DNS:        strings.Join(conf.DNSServers, ", "),
	}
	peer := WgConfigPeer{
		PersistenceKeepalive: 25,
		PublicKey:            conf.ServerKey,
		AllowedIPs:           "0.0.0.0/0, ::/0",
		Endpoint:             fmt.Sprintf("%s:%d", conf.ServerIP, conf.ServerPort),
	}

	return []byte(fmt.Sprintf("%s\n%s", iface.getText(), peer.getText()))

}

func writeFile(filename string, data []byte) error {
	err := os.WriteFile(filename, data, 0600)
	return err
}