package packet

import (
	"github.com/rxjh-emu/server/cmd/loginserver/def"
	"github.com/rxjh-emu/server/share/log"
)

var g_ServerConfig = def.ServerConfig
var g_ServerSettings = def.ServerSettings
var g_PacketHandler = def.PacketHandler
var g_RPCHandler = def.RPCHandler

// Registers network packets
func RegisterPackets() {
	log.Info("Registering packets...")

	var pk = g_PacketHandler
	pk.Register(CM_LOGIN, "CM_LOGIN", Login)
	pk.Register(CM_SERVERLIST, "CM_SERVERLIST", RequestServerList)
	pk.Register(CM_SELECT_CHANNEL, "CM_SELECT_CHANNEL", SelectChannel)

	pk.Register(SM_LOGIN, "SM_LOGIN", nil)
	pk.Register(SM_SERVERLIST, "SM_SERVERLIST", nil)
	pk.Register(SM_SELECT_CHANNEL, "SM_SELECT_CHANNEL", nil)
}

func RegisterFunc() {
	//script.RegisterFunc("sendClientPacket", sessionPacketFunc{})
	//script.RegisterFunc("sendClientMessage", clientMessageFunc{Fn: SystemMessgEx})
}
