package main

// Region data: root struct
type RegionData struct {
	Groups  Groups    `json:"groups,omitempty"`
	Regions []Regions `json:"regions,omitempty"`
}
// Region data: Group struct
type Groups struct {
	Ikev2      []GroupData      `json:"ikev2,omitempty"`
	Meta       []GroupData       `json:"meta,omitempty"`
	Ovpntcp    []GroupData    `json:"ovpntcp,omitempty"`
	Ovpnudp    []GroupData    `json:"ovpnudp,omitempty"`
	Proxysocks []GroupData `json:"proxysocks,omitempty"`
	Proxyss    []GroupData    `json:"proxyss,omitempty"`
	Wg         []GroupData         `json:"wg,omitempty"`
}

// Region data: Group.subelement struct
type GroupData struct {
	Name  string `json:"name,omitempty"`
	Ports []int  `json:"ports,omitempty"`
}

// Region data Regions.Servers.subelements struct
type ServerData struct {
	IP  string `json:"ip,omitempty"`
	Cn  string `json:"cn,omitempty"`
	Van bool   `json:"van,omitempty"`
}

// Region data Regions.Servers struct
type Servers struct {
	Ikev2   []ServerData   `json:"ikev2,omitempty"`
	Meta    []ServerData    `json:"meta,omitempty"`
	Ovpntcp []ServerData `json:"ovpntcp,omitempty"`
	Ovpnudp []ServerData `json:"ovpnudp,omitempty"`
	Wg      []ServerData      `json:"wg,omitempty"`
}

// Region data Regions struct
type Regions struct {
	ID          string  `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Country     string  `json:"country,omitempty"`
	AutoRegion  bool    `json:"auto_region,omitempty"`
	DNS         string  `json:"dns,omitempty"`
	PortForward bool    `json:"port_forward,omitempty"`
	Geo         bool    `json:"geo,omitempty"`
	Offline     bool    `json:"offline,omitempty"`
	Servers     Servers `json:"servers,omitempty"`
}

// PIA Auth Token struct
type PIAToken struct{
	Token string `json:"token"`
}