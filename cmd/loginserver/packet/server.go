package packet

import (
	"github.com/rxjh-emu/server/share/log"
	"github.com/rxjh-emu/server/share/models/server"
	"github.com/rxjh-emu/server/share/network"
	"github.com/rxjh-emu/server/share/rpc"
	"github.com/rxjh-emu/server/share/util"
	"github.com/samber/lo"
)

func RequestServerList(session *network.Session, reader *network.Reader) {
	var serverListRes server.ListRes
	g_RPCHandler.Call(rpc.ServerList, server.ListReq{}, &serverListRes)

	packet := network.NewWriter(SM_SERVERLIST)
	packet.WriteInt16(len(serverListRes.List))

	for _, server := range serverListRes.List {
		packet.WriteInt16(server.Id)
		packet.WriteString(server.Name)
		packet.WriteInt16(22)
		packet.WriteInt16(0)
		packet.WriteInt16(1) // ?

		packet.WriteInt32(len(server.List))
		for _, channel := range server.List {
			percent := (int(channel.CurrentUsers) * 100 / int(channel.MaxUsers))
			packet.WriteInt16(channel.Id)
			packet.WriteString(channel.Name)
			packet.WriteInt16(percent)
			packet.WriteInt16(21)
		}
	}

	session.Send(packet)
}

func SelectChannel(session *network.Session, reader *network.Reader) {
	serverID := reader.ReadInt16()
	channelIdx := reader.ReadInt16()

	log.Debugf("Selected Server: %d, Channel Index: %d", serverID, channelIdx)

	var listRes server.ListRes
	g_RPCHandler.Call(rpc.ServerList, server.ListReq{}, &listRes)

	selectedServer, found := lo.Find(listRes.List, func(s server.ServerItem) bool {
		return s.Id == byte(serverID)
	})
	if found {
		selectedChannel := selectedServer.List[channelIdx]

		packet := network.NewWriter(SM_SELECT_CHANNEL)
		packet.WriteString(util.ParseIPFromInt32(selectedChannel.Ip))
		packet.WriteInt16(selectedChannel.Port)
		packet.WriteString(session.Data.Username) // username
		session.Send(packet)
	}
}
