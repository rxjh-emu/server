package packet

import (
	"github.com/rxjh-emu/server/share/network"
)

func EnterLobby(session *network.Session, reader *network.Reader) {
	packet := network.NewWriter(SM_ENTER_LOBBY)
	packet.WriteInt32(631327)
	packet.WriteInt16(-1)
	packet.WriteInt16(0)

	session.Send(packet)
}
