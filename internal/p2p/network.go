package p2p

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"maps"
	"net"
	"os"
	"strings"
	"sync"
)

type Network struct {
	Addr     string
	Peers    map[net.Conn]string
	Mu       sync.Mutex
	Listener net.Listener
}

// InitPeerFromCLI initializes a Peer by collecting CLI and terminal input
func (n *Network) InitPeerFromCLI() *Peer {
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
		RemoteAddr:    n.normalizeAddr(addr),
		ShouldConnect: addr != "",
		SeenMessages:  make(map[string]bool),
	}
}

// NewPeer returns the initialized peer (placeholder for extended config)
func NewPeer(config *Peer) *Peer {
	return config
}

// StartServer begins listening for incoming TCP connections
func (n *Network) StartServer(addr string) {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	n.Listener = ln
	n.Peers = peers
	log.Printf("Listening on %s", n.Listener.Addr().String())

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Accept error:", err)
			continue
		}
		n.Mu.Lock()
		n.Peers[conn] = conn.RemoteAddr().String()
		n.Mu.Unlock()

		go n.handleConnection(conn)
	}
}

// ConnectTo connects this peer to another peer
func (n *Network) ConnectTo(addr string) {
	addr = n.normalizeAddr(addr)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Printf("Connection to %s failed: %v", addr, err)
		return
	}
	n.Mu.Lock()
	n.Peers[conn] = addr
	n.Mu.Unlock()

	// Announce self
	n.send(conn, Message{Type: "peer", From: addr})

	go n.handleConnection(conn)
}

// handleConnection handles incoming or outgoing connection
func (n *Network) handleConnection(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		msg := scanner.Text()
		fmt.Printf("[%s] %s\n", conn.RemoteAddr(), msg)
	}
}

func (n *Network) send(conn net.Conn, msg Message) error {
	jsonData, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("error here")
		return err
	}
	fmt.Fprintln(conn, string(jsonData))
	return nil
}

func (n *Network) Broadcast(msg Message) {
	n.Mu.Lock()
	peerCopy := make(map[net.Conn]string, len(peers))
	maps.Copy(peerCopy, peers)
	n.Mu.Unlock()

	data, _ := json.Marshal(msg)
	data = append(data, '\n')

	for conn, addr := range peerCopy {
		_, err := conn.Write(data)
		if err != nil {
			fmt.Printf("Failed to send to %s, removing\n", addr)
			n.Mu.Lock()
			delete(peers, conn)
			n.Mu.Unlock()
			conn.Close()
		}
	}
}

func (n *Network) forward(sender net.Conn, msg Message) {
	n.Mu.Lock()
	defer n.Mu.Unlock()
	for conn := range peers {
		if conn != sender {
			n.send(conn, msg)
		}
	}
}

func (n *Network) normalizeAddr(addr string) string {
	if !strings.Contains(addr, ":") {
		return "127.0.0.1:" + addr
	}
	if strings.HasPrefix(addr, ":") {
		return "127.0.0.1" + addr
	}
	return addr
}

func (n *Network) sendToAddr(addr string, msg Message) {
	data, _ := json.Marshal(msg)
	data = append(data, '\n')

	for conn, peerAddr := range peers {
		if peerAddr == addr {
			_, err := conn.Write(data)
			if err != nil {
				fmt.Printf("Failed to send to %s, removing\n", addr)
			}
			n.send(conn, msg)
		}
	}
}
