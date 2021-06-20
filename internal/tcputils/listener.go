package tcputils

import (
	"fmt"
	"net"
)

// LocalhostWithRandomPort contains localhost address with local port.
// See https://godoc.org/net#Listen for details.
const LocalhostWithRandomPort = "127.0.0.1:0"

const protocolTCP = "tcp"

// Listener returns a configured *net.TCPListener that listens on specified address.
func Listener(address string) (*net.TCPListener, error) {
	tcpAddr, err := net.ResolveTCPAddr(protocolTCP, address)
	if err != nil {
		return nil, fmt.Errorf("unable to setup address: %w", err)
	}

	tcpListener, err := net.ListenTCP(protocolTCP, tcpAddr)
	if err != nil {
		return nil, fmt.Errorf("unable to listen: %w", err)
	}

	return tcpListener, nil
}

// LocalListenerWithRandomPort returns a configured *net.TCPListener that listens on localhost with random port.
func LocalListenerWithRandomPort() (*net.TCPListener, error) {
	return Listener(LocalhostWithRandomPort)
}
