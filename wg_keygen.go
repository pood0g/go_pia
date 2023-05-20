package main

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/curve25519"
)

type KeyPair struct {
	prvKey string
	pubKey string
}

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
	handleFatal(err)
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

func genKeyPair() KeyPair {
	
	prvKey, pubKey := genPubkey(genPrivKey())

	prvKeyb64 := base64.StdEncoding.EncodeToString(*prvKey)
	pubKeyb64 := base64.StdEncoding.EncodeToString(*pubKey)

	return KeyPair{
		prvKey: prvKeyb64,
		pubKey: pubKeyb64,
	}
}

