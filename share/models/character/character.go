package character

import "github.com/rxjh-emu/server/share/models/world"

// Character check name status
const (
	NameInUse  = 0
	NameCanUse = 1
)

const (
	CreateCharacterFail    = 0
	CreateCharacterSuccess = 1
)

// Character class
const (
	Blademan = 1
	Swordman = 2
	Spearman = 3
	Archer   = 4
	Healer   = 5
	Ninja    = 6
	Busker   = 7
	Fighter  = 8
	Hanbi    = 9
	Arin     = 10
	Maeyujin = 11
	Noho     = 12
	Miko     = 13
)

// Character struct
type Character struct {
	Id        int    `db:"id"`
	Name      string `db:"name"`
	Slot      int    `db:"slot"`
	Class     int    `db:"class"`
	Level     int    `db:"level"`
	Exp       int    `db:"exp"`
	Ki        int    `db:"ki"`
	Spr       int    `db:"spr"`    // Basic Status
	Str       int    `db:"str"`    // Basic Status
	Stm       int    `db:"stm"`    // Basic Status
	Dex       int    `db:"dex"`    // Basic Status // Agility
	Fame      int    `db:"fame"`   // Basic Status // Honor
	Morals    int    `db:"morals"` // Basic Status // Karma
	Attack    int    // Combat Status
	Defence   int    // Combat Status
	Accuracy  int    // Combat Status
	Dodge     int    // Combat Status
	CurrentHp int
	CurrentMp int
	CurrentRp int
	MaxHp     int `db:"hp"`
	MaxMp     int `db:"mp"`
	MaxRp     int `db:"rp"`
	Gender    int `db:"gender"`
	Hair      int `db:"hair"`
	HairColor int `db:"hair_color"`
	Face      int `db:"face"`
	Voice     int `db:"voice"`

	Position world.Position
}

type ListReq struct {
	Account int
	Server  byte
}

type ListRes struct {
	List []Character
}

// Check name request struct
type CheckNameReq struct {
	Name   string
	Server byte
}

// Check name response struct
type CheckNameRes struct {
	Result int
}

// Character create request struct
type CreateCharacterReq struct {
	Server    byte
	AccountId int
	Name      string
	Class     int
	Hair      int
	HairColor int
	Face      int
	Voice     int
	Gender    int
}

type CreateCharacterRes struct {
	Result int
}
