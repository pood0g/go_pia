package main

// go_pia config variables
type goPiaConfig struct {
	PiaUser   string `json:"pia_user,omitempty"`
	PiaPass   string `json:"pia_pass,omitempty"`
	PiaRegion string `json:"pia_region,omitempty"`
	TransUser string `json:"trans_user,omitempty"`
	TransPass string `json:"trans_pass,omitempty"`
}

// Region data: root struct
type RegionData struct {
	Groups  Groups   `json:"groups,omitempty"`
	Regions []Region `json:"regions,omitempty"`
}

// Region data: Group struct
type Groups struct {
	Ikev2      []GroupData `json:"ikev2,omitempty"`
	Meta       []GroupData `json:"meta,omitempty"`
	Ovpntcp    []GroupData `json:"ovpntcp,omitempty"`
	Ovpnudp    []GroupData `json:"ovpnudp,omitempty"`
	Proxysocks []GroupData `json:"proxysocks,omitempty"`
	Proxyss    []GroupData `json:"proxyss,omitempty"`
	Wg         []GroupData `json:"wg,omitempty"`
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
	Ikev2   []ServerData `json:"ikev2,omitempty"`
	Meta    []ServerData `json:"meta,omitempty"`
	Ovpntcp []ServerData `json:"ovpntcp,omitempty"`
	Ovpnudp []ServerData `json:"ovpnudp,omitempty"`
	Wg      []ServerData `json:"wg,omitempty"`
}

// Region data Region struct
type Region struct {
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
type PIAToken struct {
	Token string `json:"token"`
}

type PIAConfig struct {
	Status     string   `json:"status"`
	ServerKey  string   `json:"server_key"`
	ServerPort int16    `json:"server_port"`
	ServerIP   string   `json:"server_ip"`
	ServerVIP  string   `json:"server_vip"`
	PeerIP     string   `json:"peer_ip"`
	PeerPubkey string   `json:"peer_pubkey"`
	DNSServers []string `json:"dns_servers"`
}

type PIAPayloadAndSignature struct {
	Status    string `json:"status"`
	Payload   string `json:"payload"`
	Signature string `json:"signature"`
}

type PIAPFPayload struct {
	Token     string `json:"token"`
	Port      uint16 `json:"port"`
	ExpiresAt string `json:"expires_at"`
}

type PIAPFStatus struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// Wireguard Keypair struct
type WGKeyPair struct {
	pubKey string
	prvKey string
}

// type AppConfig struct {
// 	username string
// 	password string
// 	region uint8
// }

type WgConfigPeer struct {
	PersistenceKeepalive uint8
	PublicKey            string
	AllowedIPs           string
	Endpoint             string
}

type WgConfigInterface struct {
	Address    string
	PrivateKey string
	DNS        string
}

type WgConfig interface {
	getText() string
}

type TransmissionSettings struct {
	AltSpeedDown                     int    `json:"alt-speed-down,omitempty"`
	AltSpeedTimeBegin                int    `json:"alt-speed-time-begin,omitempty"`
	AltSpeedTimeDay                  int    `json:"alt-speed-time-day,omitempty"`
	AltSpeedTimeEnabled              bool   `json:"alt-speed-time-enabled,omitempty"`
	AltSpeedTimeEnd                  int    `json:"alt-speed-time-end,omitempty"`
	AltSpeedUp                       int    `json:"alt-speed-up,omitempty"`
	AnnounceIP                       string `json:"announce-ip,omitempty"`
	AnnounceIPEnabled                bool   `json:"announce-ip-enabled,omitempty"`
	AntiBruteForceEnabled            bool   `json:"anti-brute-force-enabled,omitempty"`
	AntiBruteForceThreshold          int    `json:"anti-brute-force-threshold,omitempty"`
	BindAddressIpv4                  string `json:"bind-address-ipv4,omitempty"`
	BindAddressIpv6                  string `json:"bind-address-ipv6,omitempty"`
	BlocklistEnabled                 bool   `json:"blocklist-enabled,omitempty"`
	BlocklistURL                     string `json:"blocklist-url,omitempty"`
	CacheSizeMb                      int    `json:"cache-size-mb,omitempty"`
	DefaultTrackers                  string `json:"default-trackers,omitempty"`
	DhtEnabled                       bool   `json:"dht-enabled,omitempty"`
	DownloadDir                      string `json:"download-dir,omitempty"`
	DownloadQueueEnabled             bool   `json:"download-queue-enabled,omitempty"`
	DownloadQueueSize                int    `json:"download-queue-size,omitempty"`
	Encryption                       int    `json:"encryption,omitempty"`
	IdleSeedingLimit                 int    `json:"idle-seeding-limit,omitempty"`
	IdleSeedingLimitEnabled          bool   `json:"idle-seeding-limit-enabled,omitempty"`
	IncompleteDir                    string `json:"incomplete-dir,omitempty"`
	IncompleteDirEnabled             bool   `json:"incomplete-dir-enabled,omitempty"`
	LpdEnabled                       bool   `json:"lpd-enabled,omitempty"`
	MessageLevel                     int    `json:"message-level,omitempty"`
	PeerCongestionAlgorithm          string `json:"peer-congestion-algorithm,omitempty"`
	PeerIDTTLHours                   int    `json:"peer-id-ttl-hours,omitempty"`
	PeerLimitGlobal                  int    `json:"peer-limit-global,omitempty"`
	PeerLimitPerTorrent              int    `json:"peer-limit-per-torrent,omitempty"`
	PeerPort                         uint16 `json:"peer-port,omitempty"`
	PeerPortRandomHigh               int    `json:"peer-port-random-high,omitempty"`
	PeerPortRandomLow                int    `json:"peer-port-random-low,omitempty"`
	PeerPortRandomOnStart            bool   `json:"peer-port-random-on-start,omitempty"`
	PeerSocketTos                    string `json:"peer-socket-tos,omitempty"`
	PexEnabled                       bool   `json:"pex-enabled,omitempty"`
	PortForwardingEnabled            bool   `json:"port-forwarding-enabled,omitempty"`
	Preallocation                    int    `json:"preallocation,omitempty"`
	PrefetchEnabled                  bool   `json:"prefetch-enabled,omitempty"`
	QueueStalledEnabled              bool   `json:"queue-stalled-enabled,omitempty"`
	QueueStalledMinutes              int    `json:"queue-stalled-minutes,omitempty"`
	RatioLimit                       int    `json:"ratio-limit,omitempty"`
	RatioLimitEnabled                bool   `json:"ratio-limit-enabled,omitempty"`
	RenamePartialFiles               bool   `json:"rename-partial-files,omitempty"`
	RPCAuthenticationRequired        bool   `json:"rpc-authentication-required,omitempty"`
	RPCBindAddress                   string `json:"rpc-bind-address,omitempty"`
	RPCEnabled                       bool   `json:"rpc-enabled,omitempty"`
	RPCHostWhitelist                 string `json:"rpc-host-whitelist,omitempty"`
	RPCHostWhitelistEnabled          bool   `json:"rpc-host-whitelist-enabled,omitempty"`
	RPCPassword                      string `json:"rpc-password,omitempty"`
	RPCPort                          int    `json:"rpc-port,omitempty"`
	RPCSocketMode                    string `json:"rpc-socket-mode,omitempty"`
	RPCURL                           string `json:"rpc-url,omitempty"`
	RPCUsername                      string `json:"rpc-username,omitempty"`
	RPCWhitelist                     string `json:"rpc-whitelist,omitempty"`
	RPCWhitelistEnabled              bool   `json:"rpc-whitelist-enabled,omitempty"`
	ScrapePausedTorrentsEnabled      bool   `json:"scrape-paused-torrents-enabled,omitempty"`
	ScriptTorrentAddedEnabled        bool   `json:"script-torrent-added-enabled,omitempty"`
	ScriptTorrentAddedFilename       string `json:"script-torrent-added-filename,omitempty"`
	ScriptTorrentDoneEnabled         bool   `json:"script-torrent-done-enabled,omitempty"`
	ScriptTorrentDoneFilename        string `json:"script-torrent-done-filename,omitempty"`
	ScriptTorrentDoneSeedingEnabled  bool   `json:"script-torrent-done-seeding-enabled,omitempty"`
	ScriptTorrentDoneSeedingFilename string `json:"script-torrent-done-seeding-filename,omitempty"`
	SeedQueueEnabled                 bool   `json:"seed-queue-enabled,omitempty"`
	SeedQueueSize                    int    `json:"seed-queue-size,omitempty"`
	SpeedLimitDown                   int    `json:"speed-limit-down,omitempty"`
	SpeedLimitDownEnabled            bool   `json:"speed-limit-down-enabled,omitempty"`
	SpeedLimitUp                     int    `json:"speed-limit-up,omitempty"`
	SpeedLimitUpEnabled              bool   `json:"speed-limit-up-enabled,omitempty"`
	StartAddedTorrents               bool   `json:"start-added-torrents,omitempty"`
	TCPEnabled                       bool   `json:"tcp-enabled,omitempty"`
	TorrentAddedVerifyMode           string `json:"torrent-added-verify-mode,omitempty"`
	TrashOriginalTorrentFiles        bool   `json:"trash-original-torrent-files,omitempty"`
	Umask                            string `json:"umask,omitempty"`
	UploadSlotsPerTorrent            int    `json:"upload-slots-per-torrent,omitempty"`
	UtpEnabled                       bool   `json:"utp-enabled,omitempty"`
}
