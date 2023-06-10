package main

// HTTP String constants used within the application.
const REGION_URL string = "https://serverlist.piaservers.net/vpninfo/servers/v6"
const TOKEN_URL string = "https://www.privateinternetaccess.com/api/client/v2/token"
const CT_FORM string = "application/x-www-form-urlencoded"

// CA file for PIA server trust
const PIA_CERT string = "./ca.rsa.4096.crt"

// Config file
const CONFIG_FILE string = "config.json"