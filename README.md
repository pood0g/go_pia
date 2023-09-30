# go_pia

go_pia is intended to be used in a docker container running transmission connected to a Private Internet Access wireguard vpn.

Its probably easier to use the piactl tool, but I wanted to practice my golang skillz

## Building

```sh
CGO_ENABLED=0 go build --ldflags "-w -s"
```

TODO:

- Work on proper error handling 