package main

import (
	"fmt"
)

func (p WgConfigPeer) getText() string {
	config := "[Peer]\n"
	config += fmt.Sprintf("PersistentKeepalive = %d\n", p.PersistenceKeepalive)
	config += fmt.Sprintf("PublicKey = %s\n", p.PublicKey)
	config += fmt.Sprintf("AllowedIPs = %s\n", p.AllowedIPs)
	config += fmt.Sprintf("Endpoint = %s\n", p.Endpoint)
	return config
}

func (i WgConfigInterface) getText() string {
	config := "[Interface]\n"
	config += fmt.Sprintf("Address = %s\n", i.Address)
	config += fmt.Sprintf("PrivateKey = %s\n", i.PrivateKey)
	config += fmt.Sprintf("DNS = %s\n", i.DNS)
	config += fmt.Sprintf("PostUp = %s\n", i.PostUp)
	config += fmt.Sprintf("PreDown = %s\n", i.PreDown)
	return config
}
