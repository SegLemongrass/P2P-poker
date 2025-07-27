package p2p

type DiscoveryPacket struct {
	Name string `json:"name"`
	Addr string `json:"addr"` // their TCP listening address
}
