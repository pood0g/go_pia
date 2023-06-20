package main

// Terminal colours

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"
)

// HTTP String constants used within the application.
const REGION_URL string = "https://serverlist.piaservers.net/vpninfo/servers/v6"
const TOKEN_URL string = "https://www.privateinternetaccess.com/api/client/v2/token"
const CT_FORM string = "application/x-www-form-urlencoded"

// CA file for PIA server trust
const PIA_CERT string = "./ca.rsa.4096.crt"

// Config file
const CONFIG_FILE string = "config.json"
