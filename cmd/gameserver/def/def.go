package def

import (
	"os"
	"strconv"

	"github.com/rxjh-emu/server/cmd/gameserver/protocol"
	"github.com/rxjh-emu/server/share/network"
	"github.com/rxjh-emu/server/share/rpc"
)

var ServerConfig = &Config{}
var ServerSettings = &Settings{}
var NetworkManager = &network.Network{}
var PacketHandler = &network.PacketHandler{}
var RPCHandler = &rpc.Client{}
var NetworkProtocol = protocol.NewGameProtocol()

// init function, which runs before main()
func init() {
	if len(os.Args) > 2 {
		ServerSettings.ServerId = 1 // Fix to Id 1 for now
		// if id, err := strconv.Atoi(os.Args[1]); err == nil {
		// 	ServerSettings.ServerId = id
		// } else {
		// 	ServerSettings.ServerId = 1
		// }

		if id, err := strconv.Atoi(os.Args[1]); err == nil { // old os.Args[2] = ChannelId
			ServerSettings.ChannelId = id
		} else {
			ServerSettings.ChannelId = 1
		}
	} else {
		ServerSettings.ServerId = 1
		ServerSettings.ChannelId = 1
	}
}

// Returns GameServer name with id's
func GetName() string {
	// var str = "GameServer_" + strconv.Itoa(ServerSettings.ServerId)
	// str += "_" + strconv.Itoa(ServerSettings.ChannelId)
	var str = "GameServer_" + strconv.Itoa(ServerSettings.ChannelId)
	return str
}
