// internal/p2p/protocol.go
package p2p

// MessageType defines the string values used to identify message purposes.
const (
	MessageTypePeer        = "peer"
	MessageTypePeerList    = "peerlist"
	MessageTypeChat        = "chat"
	MessageTypeJoin        = "join"
	MessageTypeStartGame   = "startgame"
	MessageTypeGameState   = "gamestate"
	MessageTypePrivateHand = "privatehand"
	MessageTypePlayerList  = "playerlist"
	MessageTypeTurn        = "turn"
	MessageTypeCheck       = "check"
	MessageTypeCall        = "call"
	MessageTypeRaise       = "raise"
	MessageTypeFold        = "fold"
	MessageTypePhase       = "phase"
)

// ValidMessageTypes is a set of known valid message type strings.
var ValidMessageTypes = map[string]struct{}{
	MessageTypePeer:        {},
	MessageTypePeerList:    {},
	MessageTypeChat:        {},
	MessageTypeJoin:        {},
	MessageTypeStartGame:   {},
	MessageTypeGameState:   {},
	MessageTypePrivateHand: {},
	MessageTypePlayerList:  {},
	MessageTypeTurn:        {},
	MessageTypeCheck:       {},
	MessageTypeCall:        {},
	MessageTypeRaise:       {},
	MessageTypeFold:        {},
	MessageTypePhase:       {},
}

// IsValidMessageType returns true if the message type is recognized by the protocol.
func IsValidMessageType(t string) bool {
	_, ok := ValidMessageTypes[t]
	return ok
}
