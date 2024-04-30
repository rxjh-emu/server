package context

import (
	"github.com/rxjh-emu/server/share/network"
	"github.com/ubis/Freya/share/models/server"
)

// CellHandler defines the interface for interacting with a world map cell.
type CellHandler interface {
	GetId() (byte, byte)
	AddPlayer(session *network.Session)
	RemovePlayer(session *network.Session)
	Send(pkt *network.Writer)
}

// WorldHandler defines the interface for interacting with a game world.
type WorldHandler interface {
	EnterWorld(session *network.Session)
	ExitWorld(session *network.Session, reason server.DelUserType)
	AdjustCell(session *network.Session)
	BroadcastSessionPacket(session *network.Session, pkt *network.Writer)
	// FindWarp(warp byte) *Warp
	IsMovable(x, y int) bool
	// DropItem(item *inventory.Item, owner int32, x, y int) bool
	// PickItem(id int32) *inventory.Item
	// PeekItem(id int32, key uint16) ItemHandler
}

// WorldManagerHandler defines the interface for interacting with a world manager.
type WorldManagerHandler interface {
	FindWorld(id byte) WorldHandler
	// GetWarps(world byte) []Warp
}
