package testutils

import (
	"fmt"
	"net"
)

const protocolTCP = "tcp"

// LocalListener returns *net.TCPListener on localhost and random port.
func LocalListener() (*net.TCPListener, error) {
	addr, err := net.ResolveTCPAddr(protocolTCP, "127.0.0.1:0") // use random port: https://godoc.org/net#Listen
	if err != nil {
		return nil, fmt.Errorf("unable to setup address: %w", err)
	}

	listener, err := net.ListenTCP(protocolTCP, addr)
	if err != nil {
		return nil, fmt.Errorf("unable to listen: %w", err)
	}

	return listener, nil
}
