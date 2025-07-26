package p2p

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

// InitPeerFromCLI initializes a Peer by collecting CLI and terminal input
func InitPeerFromCLI() *Peer {
	tcpPort := flag.String("port", "9000", "TCP listening port")
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your player name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Connect to another peer? (host:port or blank): ")
	addr, _ := reader.ReadString('\n')
	addr = strings.TrimSpace(addr)

	return &Peer{
		Name:          name,
		Addr:          ":" + *tcpPort,
		RemoteAddr:    addr,
		ShouldConnect: addr != "",
		Peers:         make(map[net.Conn]string),
		SeenMessages:  make(map[string]bool),
	}
}

// NewPeer returns the initialized peer (placeholder for extended config)
func NewPeer(config *Peer) *Peer {
	return config
}

// StartServer begins listening for incoming TCP connections
func (p *Peer) StartServer() {
	ln, err := net.Listen("tcp", p.Addr)
	if err != nil {
		log.Fatalf("Failed to start TCP server: %v", err)
	}
	p.Listener = ln
	log.Printf("Listening on %s", p.Addr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Accept error:", err)
			continue
		}
		p.Peers[conn] = conn.RemoteAddr().String()
		go p.handleConnection(conn)
	}
}

// ConnectTo connects this peer to another peer
func (p *Peer) ConnectTo(addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Printf("Connection to %s failed: %v", addr, err)
		return
	}
	p.Peers[conn] = addr
	go p.handleConnection(conn)
}

// handleConnection handles incoming or outgoing connection
func (p *Peer) handleConnection(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		msg := scanner.Text()
		fmt.Printf("[%s] %s\n", conn.RemoteAddr(), msg)
	}
}
