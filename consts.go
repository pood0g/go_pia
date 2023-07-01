package main

// Terminal colours

const (
	RESET    = "\033[0m"
	RED      = "\033[31m"
	GREEN    = "\033[32m"
	YELLOW   = "\033[33m"
	LOGERROR = RED + "ERROR:" + RESET
	LOGWARN  = YELLOW + "WARN:" + RESET
	LOGINFO  = GREEN + "INFO:" + RESET
)

// HTTP String constants used within the application.
const REGION_URL string = "https://serverlist.piaservers.net/vpninfo/servers/v6"
const TOKEN_URL string = "https://www.privateinternetaccess.com/api/client/v2/token"
const CT_FORM string = "application/x-www-form-urlencoded"

// CA file for PIA server trust
const PIA_CERT string = "./ca.rsa.4096.crt"

// Config file
const CONFIG_FILE string = "go_pia_config.json"
const T_CONF_FILE string = "/config/settings.json"
const T_CONF_DIR string = "/config"
