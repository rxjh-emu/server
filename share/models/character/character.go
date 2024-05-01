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
	Id        int
	Name      string
	Slot      int
	Class     int
	Level     int
	Exp       int
	Ki        int
	Spr       int // Basic Status
	Str       int // Basic Status
	Stm       int // Basic Status
	Dex       int // Basic Status // Agility
	Fame      int // Basic Status // Honor
	Morals    int // Basic Status // Karma
	Attack    int // Combat Status
	Defence   int // Combat Status
	Accuracy  int // Combat Status
	Dodge     int // Combat Status
	CurrentHp int
	CurrentMp int
	CurrentRp int
	MaxHp     int
	MaxMp     int
	MaxRp     int
	Gender    int
	Hair      int
	HairColor int
	Face      int
	Voice     int

	Position world.Position
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
