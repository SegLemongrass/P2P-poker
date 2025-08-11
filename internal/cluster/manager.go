package cluster

import (
	"errors"
	"sync"

	"p2poker/internal/protocol"
	"p2poker/internal/table"
	"p2poker/pkg/types"
)

type TableManager struct {
	self   protocol.NodeID
	clock  *protocol.Lamport
	router *Router
	netOut chan<- protocol.NetMessage

	mu     sync.RWMutex
	tables map[protocol.TableID]*table.Table
}

func NewTableManager(self protocol.NodeID, clock *protocol.Lamport, router *Router, netOut chan<- protocol.NetMessage) *TableManager {
	return &TableManager{self: self, clock: clock, router: router, netOut: netOut, tables: make(map[protocol.TableID]*table.Table)}
}

func (m *TableManager) CreateLocalAuthorityTable(id protocol.TableID, cfg types.TableConfig) (*table.Table, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, exists := m.tables[id]; exists {
		return nil, errors.New("table exists")
	}
	in := make(chan protocol.NetMessage, 256)
	t := table.New(id, m.self, cfg, true /*authority*/, 0 /*epoch*/, m.clock, in, m.netOut)
	m.tables[id] = t
	m.router.Register(id, in)
	go t.Run()
	return t, nil
}

func (m *TableManager) AttachFollowerTable(id protocol.TableID, cfg types.TableConfig, epoch protocol.Epoch) (*table.Table, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, exists := m.tables[id]; exists {
		return nil, errors.New("table exists")
	}
	in := make(chan protocol.NetMessage, 256)
	t := table.New(id, m.self, cfg, false /*authority*/, epoch, m.clock, in, m.netOut)
	m.tables[id] = t
	m.router.Register(id, in)
	go t.Run()
	return t, nil
}

func (m *TableManager) Get(id protocol.TableID) (*table.Table, bool) {
	m.mu.RLock()
	t, ok := m.tables[id]
	m.mu.RUnlock()
	return t, ok
}
