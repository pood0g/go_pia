# go_pia

go_pia is intended to be used in a docker container running transmission connected to a Private Internet Access wireguard vpn.

Contains reverse engineered code from https://github.com/pia-foss/manual-connections

## Building

```sh
CGO_ENABLED=0 go build --ldflags "-w -s"
```

TODO:

- Work on proper error handling 
