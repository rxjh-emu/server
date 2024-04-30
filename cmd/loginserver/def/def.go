package def

import (
	"github.com/rxjh-emu/server/cmd/loginserver/protocol"
	"github.com/rxjh-emu/server/share/network"
	"github.com/rxjh-emu/server/share/rpc"
)

var ServerConfig = &Config{}
var ServerSettings = &Settings{}
var NetworkManager = &network.Network{}
var PacketHandler = &network.PacketHandler{}
var RPCHandler = &rpc.Client{}
var NetworkProtocol = protocol.NewLoginProtocol()
