package game

type WorldManager struct {
	Maps []*Map
	// Warps []struct {
	// 	Map byte
	// 	Warps []context.Warp
	// }
	// Mobs []*Mob
}

func (wm *WorldManager) Initialize() {
	wm.Maps = make([]*Map, 0)
}
