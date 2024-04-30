package rpc

import (
	"github.com/rxjh-emu/server/cmd/gameserver/def"
	"github.com/rxjh-emu/server/share/rpc"
)

var g_ServerConfig = def.ServerConfig
var g_ServerSettings = def.ServerSettings
var g_NetworkManager = def.NetworkManager
var g_RPCHandler = def.RPCHandler

func RegisterCalls() {
	g_RPCHandler.Register(rpc.UserVerify, UserVerify)
	g_RPCHandler.Register(rpc.OnlineCheck, OnlineCheck)
}
