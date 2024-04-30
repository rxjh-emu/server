package context

import (
	"sync"
)

type Context struct {
	Mutex sync.RWMutex
	// Char  *character.Character

	Cell         CellHandler
	World        WorldHandler
	WorldManager WorldManagerHandler
}
