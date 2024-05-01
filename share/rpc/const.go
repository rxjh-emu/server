package rpc

// Server related RPC's
const (
	ServerRegister = "ServerRegister"
	ServerList     = "ServerList"
)

// Account related RPC's
const (
	AuthCheck       = "AuthCheck"
	GetAccount      = "GetAccount"
	UserVerify      = "UserVerify"
	PasswdCheck     = "PasswdCheck"
	OnlineCheck     = "OnlineCheck"
	ForceDisconnect = "ForceDisconnect"
)

// Character related RPC's
const (
	LoadCharacters     = "LoadCharacters"
	CheckCharacterName = "CheckCharacterName"
	CreateCharacter    = "CreateCharacter"
	// DeleteCharacter    = "DeleteCharacter"
	// LoadCharacterData  = "LoadCharacterData"
)

// Inventory related RPC's
// const (
// 	EquipItem         = "EquipItem"
// 	UnEquipItem       = "UnEquipItem"
// 	SwapEquipmentItem = "SwapEquipmentItem"
// 	MoveEquipmentItem = "MoveEquipmentItem"
// 	AddItem           = "AddItem"
// 	StackItem         = "StackItem"
// 	RemoveItem        = "RemoveItem"
// 	SwapItem          = "SwapItem"
// 	MoveItem          = "MoveItem"
// )

// Skill related RPC's
// const (
// 	QuickLinkSet    = "QuickLinkSet"
// 	QuickLinkRemove = "QuickLinkRemove"
// 	QuickLinkSwap   = "QuickLinkSwap"
// )
