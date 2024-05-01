package world

type Position struct {
	Map int     `db:"map"`
	X   float32 `db:"x"`
	Y   float32 `db:"y"`
	Z   float32 `db:"z"`
}
