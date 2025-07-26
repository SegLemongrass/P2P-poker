package main

import (
	"p2p-poker/internal/cli"
	"p2p-poker/internal/p2p"
)

func main() {
	config := p2p.InitPeerFromCLI()
	peer := p2p.NewPeer(config)

	go peer.StartServer()

	if peer.ShouldConnect {
		peer.ConnectTo(peer.RemoteAddr)
	}

	cli.Run(peer)
}
