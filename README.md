# p2p-poker (ongoing)


Invite your friends, your cult *cough*... business men speak at golf courses... maybe it's a rainy day and need a peer-to-peer fun game to play!


```
p2p-poker/
├── cmd/
│   └── poker/                 # Entrypoint (main.go)
│       └── main.go
├── internal/
│   ├── cli/                   # Command-line interface handling
│   │   ├── commands/
│   │   │   └── peers.go
│   │   ├── command_router.go
│   │   └── cli.go
│   ├── config/                # Config parsing (env vars, flags, etc)
│   │   └── config.go
│   ├── engine/                # Orchestrates turns, betting rounds, game state transitions
│   │   ├── engine.go
│   │   └── turn.go
│   ├── game/                  # Core poker game logic (rules, hand evaluation, phases)
│   │   ├── table.go
│   │   ├── player.go
│   │   ├── hand.go
│   │   └── evaluate.go
│   └── p2p/                   # Peer-to-peer networking (TCP/UDP, discovery, messaging)
│       ├── peer.go
│       ├── network.go
│       ├── discovery.go
│       └── message.go
├── pkg/                       # Shared utilities (e.g. logger, utils)
│   └── log/                   # Logging helper
│       └── logger.go
├── assets/                    # Cards art, optional JSONs
├── go.mod
└── README.md
```

| Layer/Folder      | Responsibility                                               |
| ----------------- | ------------------------------------------------------------ |
| `cmd/`            | App entrypoint, wire dependencies                            |
| `internal/game`   | Pure poker logic: card dealing, hand ranking, blind rotation |
| `internal/p2p`    | TCP/UDP networking, peer discovery, message broadcasting     |
| `internal/engine` | Game round lifecycle: betting, advancing phase, turn control |
| `internal/cli`    | Player command handling (e.g., `/raise 50`, `/fold`)         |
| `internal/config` | Reads ports, usernames, environment setup                    |
| `pkg/`            | Non-domain utilities like logging or retry logic             |
| `assets/`         | Card definitions, presets, or future GUI art                 |
