package main

import (
	"github.com/rxjh-emu/server/cmd/gameserver/def"
	"github.com/rxjh-emu/server/cmd/gameserver/game"
	"github.com/rxjh-emu/server/cmd/gameserver/packet"
	"github.com/rxjh-emu/server/cmd/gameserver/rpc"
	"github.com/rxjh-emu/server/share/log"
	"github.com/rxjh-emu/server/share/script"
)

// globals
var g_ServerConfig = def.ServerConfig
var g_ServerSettings = def.ServerSettings
var g_NetworkManager = def.NetworkManager
var g_PacketHandler = def.PacketHandler
var g_RPCHandler = def.RPCHandler
var g_NetworkProtocol = def.NetworkProtocol

func main() {
	log.Init(def.GetName())

	// read config
	g_ServerConfig.Read()

	game := &game.WorldManager{}
	game.Initialize()

	// register events
	RegisterEvents(game)

	// register scripting engine
	script.Initialize(g_ServerConfig.ScriptDirectory)

	// register scripting functions
	packet.RegisterFunc()

	// init packet handler
	g_PacketHandler.Init()

	// register packets
	packet.RegisterPackets()

	// init RPC handler
	g_RPCHandler.Init()
	g_RPCHandler.IpAddress = g_ServerConfig.MasterIp
	g_RPCHandler.Port = g_ServerConfig.MasterPort

	// register RPC calls
	rpc.RegisterCalls()

	// start RPC handler
	g_RPCHandler.Start()

	// create network and start listening for connections
	g_NetworkManager.Init(g_ServerConfig.Port, &g_ServerSettings.Settings, g_NetworkProtocol)
}
