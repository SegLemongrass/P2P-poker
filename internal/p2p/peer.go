package p2p

import (
	"net"
	"sync"
)

// Peer represents a node in the P2P network
type Peer struct {
	Name          string
	Addr          string
	RemoteAddr    string
	ShouldConnect bool
	Listener      net.Listener
	Peers         map[net.Conn]string
	SeenMessages  map[string]bool
	SeenMu        sync.Mutex
	Ready         bool
}
