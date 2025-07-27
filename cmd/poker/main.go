package main

import (
	"p2p-poker/internal/cli"
	"p2p-poker/internal/p2p"
)

var (
	n = p2p.Network{}
)

func main() {
	config := n.InitPeerFromCLI()
	peer := p2p.NewPeer(config)
	network := p2p.Network{}

	go network.StartServer(peer.Addr)

	if peer.ShouldConnect {
		network.ConnectTo(peer.RemoteAddr)
	}

	cli.Run(peer)
}
