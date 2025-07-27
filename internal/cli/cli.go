package cli

import (
	"bufio"
	"fmt"
	"os"
	"p2p-poker/internal/p2p"
	"strings"
)

var (
	m p2p.Message
	n p2p.Network
	p p2p.Peer
)

func Run(peer *p2p.Peer) {
	fmt.Println("You can now chat. Type messages:")

	stdinScanner := bufio.NewScanner(os.Stdin)
	for stdinScanner.Scan() {
		text := strings.TrimSpace(stdinScanner.Text())

		if strings.HasPrefix(text, "/") {
			handleCommand(text)
			continue
		}

		msg := p2p.Message{
			ID:   m.GenerateID(),
			Type: "chat",
			From: peer.Name,
			Text: text,
		}

		peer.SeenMu.Lock()
		peer.SeenMessages[msg.ID] = true
		peer.SeenMu.Unlock()

		m.PrintMessage(msg)
		n.Broadcast(msg)
	}
}
