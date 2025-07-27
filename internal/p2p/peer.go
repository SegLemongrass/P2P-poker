package p2p

import (
	"net"
	"sync"
)

// Peer represents a node in the P2P network
type Peer struct {
	Name           string          `json:"name"`
	Addr           string          `json:"addr"`
	RemoteAddr     string          `json:"remote_address,omitempty"`
	ShouldConnect  bool            `json:"should_connect,omitempty"`
	Listener       net.Listener    `json:"listener"`
	Network        *Network        `json:"network"`
	SeenMessages   map[string]bool `json:"seen_messages,omitempty"`
	SeenMu         sync.Mutex      `json:"seen_mu,omitempty"`
	ReadyToReceive bool            `json:"ready_to_receive,omitempty"`
}

var peers = make(map[net.Conn]string) // conn -> listener address
