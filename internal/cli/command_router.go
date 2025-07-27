package cli

import (
	"fmt"
	"os"
	"strings"
)

func handleCommand(cmd string) {
	parts := strings.Fields(cmd)
	if len(parts) == 0 {
		return
	}

	switch parts[0] {
	case "/peers":
		fmt.Println("Known peers:")
		n.Mu.Lock()
		for _, addr := range n.Peers {
			if addr != "" {
				fmt.Println("-", addr)
			}
		}
		n.Mu.Unlock()

	case "/name":
		if len(parts) < 2 {
			fmt.Println("Usage: /name <newname>")
			return
		}
		newName := parts[1]
		fmt.Printf("Changing name from '%s' to '%s'\n", p.Name, newName)
		p.Name = newName

	case "/whoami":
		fmt.Printf("Name: %s\n", p.Name)
		fmt.Printf("Listening at: %s\n", p.Listener.Addr().String())

	case "/connect":
		if len(parts) < 2 {
			fmt.Println("Usage: /connect <ip:port>")
			return
		}
		addr := parts[1]
		go n.ConnectTo(addr)

	case "/exit":
		fmt.Println("Exiting...")
		os.Exit(0)

	case "/help":
		fmt.Println("Available commands:")
		fmt.Println("  /peers           - list known peers")
		fmt.Println("  /name <newname>  - change your name")
		fmt.Println("  /whoami          - show your name and listen address")
		fmt.Println("  /connect <ip:port> - manually connect to another peer")
		fmt.Println("  /exit            - exit the program")
		fmt.Println("  /help            - show this help menu")

	default:
		fmt.Println("Unknown command:", cmd)
	}
}
