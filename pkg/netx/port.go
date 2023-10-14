package netx

import (
	"net"
	"strconv"
)

const (
	_startPort = 30000
	_endPort   = 32767
)

// GetAvailablePort will get an available port.
func GetAvailablePort() int {
	for port := _startPort; port <= _endPort; port++ { // Ports range from 1 to 65535
		address := ":" + strconv.Itoa(port)
		listener, err := net.Listen("tcp", address)
		if err == nil {
			listener.Close()
			return port
		}
	}

	return 0 // Return 0 when no available port is found
}
