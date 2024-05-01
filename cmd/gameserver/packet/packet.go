package packet

import (
	"github.com/rxjh-emu/server/cmd/gameserver/def"
	"github.com/rxjh-emu/server/share/log"
)

var g_ServerConfig = def.ServerConfig
var g_ServerSettings = def.ServerSettings
var g_PacketHandler = def.PacketHandler
var g_RPCHandler = def.RPCHandler
var g_NetworkManager = def.NetworkManager

func RegisterPackets() {
	log.Info("Registering packets...")

	var pk = g_PacketHandler
	pk.Register(CM_AUTHORIZE, "CM_AUTHORIZE", Authorize)
	pk.Register(CM_ENTER_LOBBY, "CM_ENTER_LOBBY", EnterLobby)
	pk.Register(CM_CHARACTER_LIST, "CM_CHARACTER_LIST", CharacterList)
	pk.Register(CM_CHECK_NAME, "CM_CHECK_NAME", CheckName)
	pk.Register(CM_CHARACTER_CREATE, "CM_CHARACTER_CREATE", CreateCharacter)

	pk.Register(SM_AUTHORIZE, "SM_AUTHORIZE", nil)
	pk.Register(SM_ENTER_LOBBY, "SM_ENTER_LOBBY", nil)
	pk.Register(SM_CHARACTER_LIST, "SM_CHARACTER_LIST", nil)
	pk.Register(SM_CHECK_NAME, "SM_CHECK_NAME", nil)
	pk.Register(SM_CHARACTER_CREATE, "SM_CHARACTER_CREATE", nil)
}

func RegisterFunc() {
	// script.RegisterFunc("sendClientPacket", sessionPacketFunc{})
	// script.RegisterFunc("sendClientMessage", clientMessageFunc{Fn: SendMessage})

	// script.RegisterFunc("getPlayerLevel", playerGetLevelFunc{Fn: GetPlayerLevel})
	// script.RegisterFunc("setPlayerLevel", playerSetLevelFunc{Fn: SetPlayerLevel})
	// script.RegisterFunc("getPlayerPosition", playerPositionFunc{})
	// script.RegisterFunc("dropItem", playerDropItemFunc{})
}
