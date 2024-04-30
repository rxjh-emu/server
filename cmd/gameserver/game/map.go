package game

import "time"

type Map struct {
	Id byte

	// mobs
	mobsMove *time.Ticker
	// mobs     map[int]*Mob

	// tickers
	itemTicker *time.Ticker

	// data
	// Warps []context.Warp
}
