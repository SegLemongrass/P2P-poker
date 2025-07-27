package p2p

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

type Message struct {
	ID    string `json:"id"`
	Type  string `json:"type"` // "peer", "peerlist", "chat"
	From  string `json:"from,omitempty"`
	Text  string `json:"text,omitempty"`
	Peers []Peer `json:"peers,omitempty"`
	Addr  string `json:"addr,omitempty"`
}

// PrintMessage prints a receiving message in the terminal
func (m *Message) PrintMessage(msg Message) {
	switch msg.Type {
	case "chat":
		fmt.Printf("[%s]: %s\n", msg.From, msg.Text)
	case "system":
		fmt.Printf("[SYSTEM]: %s\n", msg.Text)
	}
}

// generateID generates a unique ID for each message
func (m *Message) GenerateID() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

// shouldSkipMessage returns whether peer is ready to receive
func (m *Message) shouldSkipMessage(reason string, peer *Peer) bool {
	if !peer.ReadyToReceive {
		fmt.Println("Skipping message:", reason)
		return true
	}
	return false
}
