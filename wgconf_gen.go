package main

import (
	"fmt"
	"strings"
)

type WgConfigPeer struct {
	PersistenceKeepalive uint8
	PublicKey string
	AllowedIPs string
	Endpoint string
}

type WgConfigInterface struct {
	Address string
	PrivateKey string
	DNS string
}

type WgConfig interface {
	getText() string
}

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
	return config
}

func genConfig(conf PIAConfig) (string) {

	iface := WgConfigInterface {
		Address: conf.PeerIP,
		PrivateKey: conf.ServerKey,
		DNS: strings.Join(conf.DNSServers, ", "),
	}
	peer := WgConfigPeer {
		PersistenceKeepalive: 25,
		PublicKey: conf.PeerPubkey,
		AllowedIPs: "0.0.0.0/0, ::/0",
		Endpoint: fmt.Sprintf("%s:%d", conf.ServerIP, conf.ServerPort),
	}

	return fmt.Sprintf("%s\n%s", iface.getText(), peer.getText())

}